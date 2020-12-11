// Package backend implements the backends for the Terraform registry
package backend

import (
	"sync"
	"time"
)

// Type represents the type of backend
type Type int

const (
	// Unknown is an unknown backend
	Unknown Type = iota
	// ProviderBackend is a provider backend
	ProviderBackend
)

// Provider defines the interface for a provider backend
type Provider interface {
	// :namespace :type
	List(string, string) (*ListResponse, error)
	// :namespace :type :version :os :arch
	Find(string, string, string, string, string) (*FindResponse, error)
}

// Module defines the interface for a module backend
type Module interface {
}

// nolint: gochecknoglobals
var (
	backendsMu sync.RWMutex
	provider   = make(map[string]Provider)
	module     = make(map[string]Module)
)

// nowFunc returns the current time; it's overridden in tests.
var nowFunc = time.Now

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, backend interface{}) {
	backendsMu.Lock()
	defer backendsMu.Unlock()

	if backend == nil {
		panic("backend: Register backend is nil")
	}

	switch b := backend.(type) {
	case Provider:
		registerProviderBackend(name, b)
	case Module:
		registerModuleBackend(name, b)
	}
}

func registerProviderBackend(name string, backend Provider) {
	if _, dup := provider[name]; dup {
		panic("backend: Register called twice for provider backend " + name)
	}

	provider[name] = backend
}

func registerModuleBackend(name string, backend Module) {
	if _, dup := module[name]; dup {
		panic("backend: Register called twice for module backend " + name)
	}

	module[name] = backend

}
