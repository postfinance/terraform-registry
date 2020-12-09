// Package registry implements the registry
package registry

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/marcsauter/terraform-registry/internal/registry/provider"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	httpStopTimeout = 10 * time.Second
)

// Registry represents the Terraform registry
type Registry struct {
	listenAddr      string
	providerBackend provider.Backend

	router   *chi.Mux
	services map[string]string
	server   *http.Server
	wg       *sync.WaitGroup

	l   *zap.SugaredLogger
	reg prometheus.Registerer
}

// New initializes the Registry.
func New(l *zap.SugaredLogger, reg prometheus.Registerer, options ...Option) (*Registry, error) {
	r := &Registry{
		l:   l,
		reg: reg,

		wg:       &sync.WaitGroup{},
		router:   chi.NewRouter(),
		services: make(map[string]string),
	}

	for _, opt := range options {
		opt(r)
	}

	r.router.Get("/.well-known/terraform.json", r.discovery)
	r.router.Get("/healthz", r.healthz)

	promHandler := func() http.Handler {
		r.reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
		r.reg.MustRegister(prometheus.NewGoCollector())
		g, _ := r.reg.(prometheus.Gatherer)

		return promhttp.InstrumentMetricHandler(
			r.reg, promhttp.HandlerFor(g, promhttp.HandlerOpts{}),
		)
	}

	r.router.Method("GET", "/metrics", promHandler())

	r.services["providers.v1"] = "/v1/providers"
	r.router.Mount("/v1/providers", provider.New(r.providerBackend).Routes())

	return r, nil
}

// Run the Terraform registry
func (reg *Registry) Run(ctx context.Context) error {
	reg.wg.Add(1)

	errChan := make(chan error)

	go func() {
		if err := reg.start(); err != nil {
			errChan <- err
		}
	}()

	for {
		select {
		case err := <-errChan:
			return err
		case <-ctx.Done():
			err := reg.stop()
			reg.wg.Wait()

			return err
		}
	}
}

func (reg *Registry) start() error {
	reg.l.Infow("starting registry server")
	reg.server = &http.Server{
		Addr:    reg.listenAddr,
		Handler: reg.router,
	}

	if err := reg.server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (reg *Registry) stop() error {
	defer reg.wg.Done()
	defer reg.l.Info("registry server stopped")

	if reg.server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), httpStopTimeout)
	defer cancel()

	return reg.server.Shutdown(ctx)
}

// discovery is the handler to serve Terraform service discovery
func (reg *Registry) discovery(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, reg.services)
}

func (reg *Registry) healthz(w http.ResponseWriter, r *http.Request) {
	status := struct{}{}

	render.JSON(w, r, status)
}
