package registry_test

import (
	"encoding/json"
	"testing"

	"github.com/marcsauter/tfregistry/internal/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArch(t *testing.T) {
	type arch struct {
		Arch registry.Arch `json:"arch"`
	}

	obj := arch{
		Arch: registry.ArchAMD64,
	}

	str := []byte(`{"arch":"amd64"}`)

	t.Run("marshal", func(t *testing.T) {
		act, err := json.Marshal(&obj)
		require.NoError(t, err)
		assert.Equal(t, str, act)
	})

	t.Run("unmarshal", func(t *testing.T) {
		act := arch{}
		err := json.Unmarshal(str, &act)
		require.NoError(t, err)
		assert.Equal(t, obj, act)
	})
}

func TestUnmarshalArch(t *testing.T) {

}
