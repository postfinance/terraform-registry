// Package provider contains the provider backend interface and all necessary types
package provider

// Backend is the interface for provider backends
// https://www.terraform.io/docs/internals/provider-registry-protocol.html
type Backend interface {
	Versions(*VersionsRequest) (*VersionsResponse, error)
	Download(*DownloadRequest) (*DownloadResponse, error)
}

// VersionsRequest implements the provider versions request
type VersionsRequest struct {
	Namespace string
	Type      string
}

// VersionsResponse implements the provider versions response
type VersionsResponse struct {
	Versions []Version `json:"versions"`
}

// Version implements the provider version information
type Version struct {
	Version   string     `json:"version"`
	Protocols []string   `json:"protocols,omitempty"`
	Platforms []Platform `json:"platforms,omitempty"`
}

// Platform implments the target platform
type Platform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

// DownloadRequest implements the provider download request
type DownloadRequest struct {
	Namespace string
	Type      string
	Version   string
	OS        string
	Arch      string
}

// DownloadResponse implements the provider download response
type DownloadResponse struct {
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
	TrustSignature string `json:"trust_signature,omitempty"`
	Source         string `json:"source,omitempty"`
	SourceURL      string `json:"source_url,omitempty"`
}
