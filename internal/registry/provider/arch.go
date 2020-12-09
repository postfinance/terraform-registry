package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Arch represents the architecture
type Arch int

// Architectures
const (
	ArchNoSet Arch = iota // 0
	ArchAMD64
	ArchARM
)

func (a Arch) String() string {
	switch a {
	case ArchAMD64:
		return "amd64"
	case ArchARM:
		return "arm"
	default:
		return ""
	}
}

// MarshalJSON implements the json.Marshaler interface
func (a Arch) MarshalJSON() ([]byte, error) {
	return bytes.NewBufferString(fmt.Sprintf("%q", a.String())).Bytes(), nil
}

// UnmarshalJSON implements the json.UnmarshalJSON interface
func (a *Arch) UnmarshalJSON(b []byte) error {
	var as string

	if err := json.Unmarshal(b, &as); err != nil {
		return err
	}

	switch strings.ToLower(as) {
	case "amd64":
		*a = ArchAMD64
	case "arm":
		*a = ArchARM
	default:
		*a = ArchNoSet
	}

	return nil
}
