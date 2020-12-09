// Package cmd implements the command line interface for the terraform-registry
package cmd

import "github.com/marcsauter/terraform-registry/internal/version"

// CLI is the Server command.
type CLI struct {
	Run runCmd `cmd:"true" help:"Start the terraform-registry." default:"true" hidden:"true"`
	Globals
}

// Globals are the global parameters for the lslb server.
type Globals struct {
	Debug      bool         `help:"Log debug output."`
	Version    version.Flag `help:"Show version information."`
	HTTPListen string       `help:"HTTP Port." default:":8080"`
}
