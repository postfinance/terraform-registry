// Package module implements the module registry as in https://www.terraform.io/docs/internals/module-registry-protocol.html
package module

import (
	"errors"
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

// errorResponse builds the error response
func errorResponse(err error) interface{} {
	return struct {
		Errors []string `json:"errors"`
	}{
		Errors: []string{err.Error()},
	}
}

// versionsHandler handles GET :namespace/:name/:provider/versions
func (a API) versionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = a.r.JSON(w, http.StatusNotImplemented, errorResponse(errors.New("not implemented")))
	}
}

// downloadHandler handles GET :namespace/:name/:provider/:version/download
func (a API) downloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = a.r.JSON(w, http.StatusNotImplemented, errorResponse(errors.New("not implemented")))
	}
}
