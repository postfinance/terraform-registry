// Package cmd implements the command line interface for the terraform-registry
package cmd

import (
	"os"
	"time"

	"github.com/marcsauter/terraform-registry/internal/version"
	"github.com/marcsauter/terraform-registry/pkg/module"
	"github.com/marcsauter/terraform-registry/pkg/provider"
	"github.com/postfinance/profiler"
)

// CLI is the Server command.
type CLI struct {
	Run runCmd `cmd:"true" help:"Start the terraform-registry." default:"true" hidden:"true"`
	Globals
}

// Globals are the global parameters for the lslb server.
type Globals struct {
	Debug           bool                 `help:"Log debug output."`
	Version         version.Flag         `help:"Show version information."`
	HTTPListen      string               `help:"HTTP Port." default:":8080"`
	ModuleBackend   moduleBackendFlags   `embed:"true" prefix:"module-backend-"`
	ProviderBackend providerBackendFlags `embed:"true" prefix:"provider-backend-"`
	Profiler        profilerFlags        `embed:"true" prefix:"profiler-"`
}

type moduleBackendFlags struct {
	URL      string `help:"URL of the module backend. If not set, the module registry will not be provided."`
	Username string `help:"Username"`
	Password string `help:"Password"`
}

func (m *moduleBackendFlags) backend() (module.Backend, error) {
	return nil, nil
}

type providerBackendFlags struct {
	URL      string `help:"URL of the provider backend. If not set, the provider registry will not be provided."`
	Username string `help:"Username"`
	Password string `help:"Password"`
}

func (p *providerBackendFlags) backend() (provider.Backend, error) {
	return nil, nil
}

type profilerFlags struct {
	Enable  bool          `help:"Enable the profiler."`
	Listen  string        `help:"The profiles listen address." default:":6666"`
	Timeout time.Duration `help:"The timeout after the pprof handler will be shutdown." default:"5m"`
}

func (p profilerFlags) New(s os.Signal, h ...profiler.Hooker) *profiler.Profiler {
	return profiler.New(
		profiler.WithAddress(p.Listen),
		profiler.WithTimeout(p.Timeout),
		profiler.WithSignal(s),
		profiler.WithHooks(h...),
	)
}
