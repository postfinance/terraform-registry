package provider_test

import (
	"encoding/json"
	"testing"

	"github.com/marcsauter/terraform-registry/pkg/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalListResponse(t *testing.T) {
	r := provider.VersionsResponse{
		Versions: []provider.Version{
			{
				Version:   "0.0.1",
				Protocols: []string{"4", "5.1"},
				Platforms: []provider.Platform{
					{
						OS:   "linux",
						Arch: "amd64",
					},
				},
			},
		},
	}

	act, err := json.Marshal(&r)
	require.NoError(t, err)

	exp := []byte(`{"versions":[{"version":"0.0.1","protocols":["4","5.1"],"platforms":[{"os":"linux","arch":"amd64"}]}]}`)

	assert.Equal(t, exp, act)
}
