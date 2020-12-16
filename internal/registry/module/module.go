// Package module implements the module registry as in https://www.terraform.io/docs/internals/module-registry-protocol.html
package module

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/marcsauter/terraform-registry/pkg/module"
	"github.com/unrolled/render"
)

// API implements the provider registry API
type API struct {
	b module.Backend
	r *render.Render
}

// New returns a new provider API
func New(b module.Backend) *API {
	return &API{
		b: b,
		r: render.New(),
	}
}

// Routes returns a router for the provider registry API
func (a API) Routes() chi.Router {
	api := chi.NewRouter()

	// :namespace/:name/:provider/versions
	// e.g. hashicorp/consul/aws/versions
	api.Get("/{namespace}/{name}/{provider}/versions", a.versionsHandler())

	// :namespace/:name/:provider/:version/download
	// e.g. hashicorp/consul/aws/0.0.1/download
	api.Get("/{namespace}/{name}/{provider}/{version}/download", a.downloadHandler())

	return api
}

func (a API) versionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func (a API) downloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
