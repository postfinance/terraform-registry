package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/marcsauter/terraform-registry/internal/cmd"
	"github.com/marcsauter/terraform-registry/internal/logger"
	ver "github.com/marcsauter/terraform-registry/internal/version"
)

//nolint:gochecknoglobals
var (
	version = "0.0.0"
	commit  = "12345678"
	date    string
)

func main() {
	v := ver.New(version, commit, date)

	cli := cmd.CLI{}

	app := kong.Parse(&cli,
		kong.Name("terraform-registry"),
		kong.Description("implements the provider registry protocol"),
		kong.Vars{
			ver.Key: v.String(),
		},
	)

	l, err := logger.New(cli.Debug)
	if err != nil {
		panic(fmt.Sprintf("could not create logger: %s", err))
	}

	if err := app.Run(&cli.Globals, l, v); err != nil {
		l.Fatal(err)
	}
}
