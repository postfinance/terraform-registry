package artifactory

import "github.com/marcsauter/terraform-registry/internal/registry/provider"

// Server represents the HTTP file server
type Server struct {
	BaseURL string
}

// New return a new HTTP file server provider.Backend
func New() (*Server, error) {
	return nil, nil
}

var _ provider.Backend = &Server{}

// List implements provider.Backend
func (s Server) List(namespace, providerType string) (*provider.ListResponse, error) {
	return nil, nil
}

// Find implements provider.Backend
func (s Server) Find(namespace, providerType, version, os, arch string) (*provider.FindResponse, error) {
	return nil, nil
}
