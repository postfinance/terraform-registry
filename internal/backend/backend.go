// Package backend implements the backends for the Terraform registry
package backend

import (
	"fmt"

	"github.com/marcsauter/terraform-registry/internal/registry/provider"
)

func Provider(name string) (provider.Backend, error) {
	switch name {
	case "fileserver":
		return nil, nil
	default:
		return nil, fmt.Errorf("unknown provider backend: %s", name)
	}
}
