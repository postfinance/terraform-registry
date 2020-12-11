// Package artifactory implements the provider.Backend interface for Artifactory
package artifactory

import (
	"github.com/marcsauter/terraform-registry/internal/artifactory"
	"github.com/marcsauter/terraform-registry/internal/backend"
	"github.com/marcsauter/terraform-registry/internal/registry/provider"
)

func init() {
	backend.Register(backend.Provider, "artifactory", &Backend{})
}

// Backend represents the Artifactory backend
type Backend struct {
	client *artifactory.Client
}

// New return a new HTTP file server provider.Backend
func New(c *artifactory.Client) (*Backend, error) {
	return nil, nil
}

var _ provider.Backend = &Backend{}

// List implements provider.Backend
func (s Backend) List(namespace, providerType string) (*provider.ListResponse, error) {
	return nil, nil
}

// Find implements provider.Backend
func (s Backend) Find(namespace, providerType, version, os, arch string) (*provider.FindResponse, error) {
	return nil, nil
}
