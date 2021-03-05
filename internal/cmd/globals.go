// Package cmd implements the command line interface for the terraform-registry
package cmd

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/postfinance/profiler"
	"github.com/postfinance/terraform-registry/pkg/artifactory"
	"github.com/postfinance/terraform-registry/pkg/module"
	"github.com/postfinance/terraform-registry/pkg/provider"
	"github.com/zbindenren/king"
)

// CLI is the Server command.
type CLI struct {
	Run runCmd `cmd:"true" help:"Start the registry." default:"true" hidden:"true"`
	Globals
}

// Globals are the global parameters for the terraform-registry.
type Globals struct {
	Debug           bool                 `help:"Log debug output."`
	Version         king.VersionFlag     `help:"Show version information"`
	HTTPListen      string               `help:"HTTP Port." default:":8080"`
	Namespace       string               `help:"The namespace portion of the address of the requested provider." required:""`
	ModuleBackend   moduleBackendFlags   `embed:"true" prefix:"module-backend-"`
	ProviderBackend providerBackendFlags `embed:"true" prefix:"provider-backend-"`
	Profiler        profilerFlags        `embed:"true" prefix:"profiler-"`
}

type moduleBackendFlags struct {
	URL      string `help:"URL of the module backend. If not set, the module registry will not be provided."`
	Username string `help:"Username"`
	Password string `help:"Password"`
}

func (m *moduleBackendFlags) backend() (module.Backend, error) {
	return nil, nil
}

type providerBackendFlags struct {
	URL               string        `help:"URL of the provider backend. If not set, the provider registry will not be provided."`
	Username          string        `help:"Username"`
	Password          string        `help:"Password"`
	CACert            string        `help:"CA certificate file." type:"existingfile"`
	GPGPublicKeyFiles []string      `help:"File(s) with ASCII armored GPG public keys" type:"existingfile" placeholder:"KEYFILE" required:""`
	Timeout           time.Duration `help:"The request timeout." default:"10s"`
}

func (p *providerBackendFlags) backend() (provider.Backend, error) {
	c := &http.Client{
		Timeout: p.Timeout,
	}

	if p.CACert != "" {
		caCert, err := ioutil.ReadFile(p.CACert)
		if err != nil {
			log.Fatal(err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)

		c.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
				RootCAs:    caCertPool,
			},
		}
	}

	return artifactory.New(c, p.URL, p.Username, p.Password, p.GPGPublicKeyFiles)
}

type profilerFlags struct {
	Enable  bool          `help:"Enable the profiler."`
	Listen  string        `help:"The profiles listen address." default:":6666"`
	Timeout time.Duration `help:"The timeout after the pprof handler will be shutdown." default:"5m"`
}

func (p profilerFlags) New(s os.Signal, h ...profiler.Hooker) *profiler.Profiler {
	return profiler.New(
		profiler.WithAddress(p.Listen),
		profiler.WithTimeout(p.Timeout),
		profiler.WithSignal(s),
		profiler.WithHooks(h...),
	)
}
