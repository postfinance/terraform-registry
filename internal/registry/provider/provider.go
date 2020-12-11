// Package provider implements the provider registry as in https://www.terraform.io/docs/internals/provider-registry-protocol.html
package provider

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

// API implements the provider registry API
type API struct {
	b Backend
	r *render.Render
}

// New returns a new provider API
func New(backend Backend) *API {
	return &API{
		b: backend,
		r: render.New(),
	}
}

// Routes returns a router for the provider registry API
func (a API) Routes() chi.Router {
	api := chi.NewRouter()

	// :namespace/:type/versions
	api.Get("/{name}/{type}/versions", a.listHandler())

	// :namespace/:type/:version/download/:os/:arch
	api.Get("/{namespace}/{type}/{version}/download/{os}/{arch}", a.findHandler())

	return api
}

func errorResponse(err error) interface{} {
	return struct {
		Errors []string `json:"errors"`
	}{
		Errors: []string{err.Error()},
	}
}

func (a API) listHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := chi.URLParam(r, "namespace")
		t := chi.URLParam(r, "type")

		resp, err := a.b.List(n, t)
		if err != nil {
			_ = a.r.JSON(w, http.StatusNotFound, errorResponse(err))

			return
		}

		_ = a.r.JSON(w, http.StatusOK, resp)
	}
}

func (a API) findHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := chi.URLParam(r, "namespace")
		t := chi.URLParam(r, "type")
		v := chi.URLParam(r, "version")
		os := chi.URLParam(r, "os")
		arch := chi.URLParam(r, "a")

		resp, err := a.b.Find(n, t, v, os, arch)
		if err != nil {
			_ = a.r.JSON(w, http.StatusNotFound, errorResponse(err))

			return
		}

		_ = a.r.JSON(w, http.StatusOK, resp)
	}
}

// ListResponse implements the response to a list request
type ListResponse struct {
	Versions []Provider `json:"versions"`
}

// Provider implements provider informations
type Provider struct {
	Version   string     `json:"version"`
	Protocols []string   `json:"protocols"`
	Platforms []Platform `json:"platforms"`
}

// Platform implements platform informations
type Platform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

// FindResponse represents the response to a find request
type FindResponse struct {
	Protocols           []string    `json:"protocols"`
	OS                  string      `json:"os"`
	Arch                string      `json:"arch"`
	Filename            string      `json:"filename"`
	DownloadURL         string      `json:"download_url"`
	ShasumsURL          string      `json:"shasums_url"`
	ShasumsSignatureURL string      `json:"shasums_signature_url"`
	Shasum              string      `json:"shasum"`
	SigningKeys         SigningKeys `json:"signing_keys"`
}

// SigningKeys represents the signing keys for providers
type SigningKeys struct {
	GPGPublicKeys []GPGPublicKey `json:"gpg_public_keys"`
}

// GPGPublicKey represents a GPG public key
type GPGPublicKey struct {
	KeyID          string `json:"key_id"`
	ASCIIArmor     string `json:"ascii_armor"`
	TrustSignature string `json:"trust_signature"`
	Source         string `json:"source"`
	SourceURL      string `json:"source_url"`
}
