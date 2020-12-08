// Package provider implements the provider registry
package provider

import "github.com/marcsauter/tfregistry/internal/registry"

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
	OS   registry.OS   `json:"os"`
	Arch registry.Arch `json:"arch"`
}
