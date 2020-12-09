// Package provider implements the provider registry as in https://www.terraform.io/docs/internals/provider-registry-protocol.html
package provider

import (
	"github.com/go-chi/chi"
)

// API implements the provider registry API
type API struct{}

// Routes returns a router for the provider registry API
func (a API) Routes() chi.Router {
	api := chi.NewRouter()

	return api
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
	OS   OS   `json:"os"`
	Arch Arch `json:"arch"`
}

// FindResponse represents the response to a find request
type FindResponse struct {
	Protocols           []string    `json:"protocols"`
	OS                  OS          `json:"os"`
	Arch                Arch        `json:"arch"`
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
