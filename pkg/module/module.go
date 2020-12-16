// Package module contains the provider backend interface and all necessary types
package module

import "net/url"

// Backend is the interface for module backends
// https://www.terraform.io/docs/internals/module-registry-protocol.html
type Backend interface {
	Versions(*VersionsRequest) (*VersionsResponse, error)
	Download(*DownloadRequest) (*DownloadResponse, error)
}

// VersionsRequest implements the versions request for available module versions
type VersionsRequest struct {
	Namespace string
	Name      string
	Provider  string
}

// VersionsResponse implements the versions response of available module version
type VersionsResponse struct {
	Modules []Module `json:"modules"`
}

// Module implements the available module versions
type Module struct {
	Versions string `json:"versions"`
}

// Version implements a specific version
type Version struct {
	Version string `json:"version"`
}

// DownloadRequest implements the download request for a specific module version
type DownloadRequest struct {
	Namespace string
	Name      string
	Provider  string
	Version   string
}

// DownloadResponse implements the download response of a specific module version
type DownloadResponse struct {
	URL url.URL
}
