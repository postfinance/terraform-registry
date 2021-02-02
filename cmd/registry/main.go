package main

import (
	"github.com/alecthomas/kong"
	"github.com/marcsauter/terraform-registry/internal/cmd"
	"github.com/postfinance/flash"
	"github.com/zbindenren/king"
)

//nolint:gochecknoglobals // there is no other way to initialize these values
var (
	version = "0.0.0"
	commit  = "12345678"
	date    string
)

func main() {
	l := flash.New(flash.WithColor())

	b, err := king.NewBuildInfo(version,
		king.WithDateString(date),
		king.WithRevision(commit),
		king.WithLocation("Europe/Zurich"),
	)
	if err != nil {
		l.Fatal(err)
	}

	cli := cmd.CLI{}
	app := kong.Parse(&cli, king.DefaultOptions(
		king.Config{
			Name:        "registry",
			Description: "Implements HashiCorp's Provider Registry Protocol for Terraform (see: https://www.terraform.io/docs/internals/provider-registry-protocol.html)",
			BuildInfo:   b,
		},
	)...)

	l.SetDebug(cli.Debug)

	if err := app.Run(&cli.Globals, l.Get()); err != nil {
		l.Fatal(err)
	}
}
