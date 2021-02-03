// Package provider implements the provider registry as in https://www.terraform.io/docs/internals/provider-registry-protocol.html
package provider

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/postfinance/terraform-registry/pkg/provider"
	"github.com/unrolled/render"
)

// API implements the provider registry API
type API struct {
	b provider.Backend
	r *render.Render
}

// New returns a new provider API
func New(b provider.Backend) *API {
	return &API{
		b: b,
		r: render.New(),
	}
}

// Routes returns a router for the provider registry API
func (a API) Routes() chi.Router {
	api := chi.NewRouter()

	// :namespace/:type/versions
	api.Get("/{namespace}/{type}/versions", a.versionsHandler())

	// :namespace/:type/:version/download/:os/:arch
	api.Get("/{namespace}/{type}/{version}/download/{os}/{arch}", a.downloadHandler())

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

// versionsHandler handles GET :namespace/:type/versions
// returns json
func (a API) versionsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := a.b.Versions(&provider.VersionsRequest{
			Namespace: chi.URLParam(r, "namespace"),
			Type:      chi.URLParam(r, "type"),
		})
		if err != nil {
			_ = a.r.JSON(w, http.StatusNotFound, errorResponse(err))

			return
		}

		_ = a.r.JSON(w, http.StatusOK, resp)
	}
}

// downloadHandler handles GET :namespace/:type/:version/download/:os/:arch
// returns json
func (a API) downloadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := a.b.Download(&provider.DownloadRequest{
			Namespace: chi.URLParam(r, "namespace"),
			Type:      chi.URLParam(r, "type"),
			Version:   chi.URLParam(r, "version"),
			OS:        chi.URLParam(r, "os"),
			Arch:      chi.URLParam(r, "arch"),
		})
		if err != nil {
			_ = a.r.JSON(w, http.StatusNotFound, errorResponse(err))

			return
		}

		_ = a.r.JSON(w, http.StatusOK, resp)
	}
}
