// Package version builds the version information
package version

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/alecthomas/kong"
)

const (
	// Key used in kong.Vars
	Key = "VersionInfo"

	versionInfotmpl = `
{{.Program}}, version {{.Version}} (revision: {{.Revision}})
  build date:       {{.BuildDate}}
  go version:       {{.GoVersion}}
`
)

// Version represents the version information
type Version struct {
	Program   string `json:"program"`
	Version   string `json:"version"`
	Revision  string `json:"revision"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
}

// New creates a new Version.
func New(version, commit, date string) *Version {
	revision := "12345678"
	if len(commit) >= 8 {
		revision = commit[:8]
	}

	program := filepath.Base(os.Args[0])

	return &Version{
		Program:   program,
		Version:   version,
		Revision:  revision,
		BuildDate: date,
		GoVersion: runtime.Version(),
	}
}

func (v *Version) String() string {
	t := template.Must(template.New("version").Parse(versionInfotmpl))

	var buf bytes.Buffer
	if err := t.ExecuteTemplate(&buf, "version", v); err != nil {
		panic(err)
	}

	return strings.TrimSpace(buf.String())
}

type Flag bool

// BeforeApply is the actual version command.
func (f Flag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Fprintln(app.Stdout, vars[Key])
	app.Exit(0)

	return nil
}
