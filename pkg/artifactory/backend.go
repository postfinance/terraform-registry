package artifactory

import (
	"github.com/marcsauter/terraform-registry/pkg/provider"
)

// Providers implements provider.Backend for Artifactory
type Providers struct {
	BaseURL string
}

// New return a new provider.Backend for Artifactory
func New() (*Providers, error) {
	return nil, nil
}

var _ provider.Backend = &Providers{}

// Versions implements provider.Backend
func (s Providers) Versions(req *provider.VersionsRequest) (*provider.VersionsResponse, error) {
	return nil, nil
}

// Download implements provider.Backend
func (s Providers) Download(req *provider.DownloadRequest) (*provider.DownloadResponse, error) {
	return nil, nil
}
