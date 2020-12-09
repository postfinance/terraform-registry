package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// OS represents the operating system
type OS int

// Operating systems
const (
	OSLinux OS = iota
	OSDarwin
	OSWindows
)

func (o OS) String() string {
	switch o {
	case OSLinux:
		return "linux"
	case OSDarwin:
		return "darwin"
	case OSWindows:
		return "windows"
	default:
		return ""
	}
}

// MarshalJSON implements the json.Marshaler interface
func (o OS) MarshalJSON() ([]byte, error) {
	return bytes.NewBufferString(fmt.Sprintf("%q", o.String())).Bytes(), nil
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (o *OS) UnmarshalJSON(b []byte) error {
	var os string

	if err := json.Unmarshal(b, &os); err != nil {
		return err
	}

	switch strings.ToLower(os) {
	case "linux":
		*o = OSLinux
	case "aarm":
		*o = OSDarwin
	default:
		*o = OSWindows
	}

	return nil
}
