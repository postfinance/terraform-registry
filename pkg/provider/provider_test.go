package provider_test

import (
	"encoding/json"
	"testing"

	"github.com/marcsauter/terraform-registry/pkg/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test marshalling of Version - required field only
func TestMarshalVersion(t *testing.T) {
	v := provider.Version{
		Version: "0.0.1",
	}

	act, err := json.Marshal(&v)
	require.NoError(t, err)

	exp := []byte(`{"version":"0.0.1"}`)

	assert.Equal(t, exp, act)
}

// Test marshalling of VersionResponse
func TestMarshalVersionResponse(t *testing.T) {
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

// Test marshalling of DownloadResponse
func TestMarshalDownloadResponse(t *testing.T) {
	r := provider.DownloadResponse{
		Protocols:           []string{"5.0"},
		OS:                  "linux",
		Arch:                "amd64",
		Filename:            "terraform-provider-example_linux_x86_64-1.1.8.zip",
		DownloadURL:         "https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_linux_x86_64-1.1.8.zip",
		ShasumsURL:          "https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_1.1.8_SHA256SUMS.txt",
		ShasumsSignatureURL: "https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_1.1.8_SHA256SUMS.txt.sig",
		Shasum:              "d7dddb0a94c4388e4e3bf5f68faea18c46eab8aaceaec8954b269a4a29f13c29",
		SigningKeys: provider.SigningKeys{
			GPGPublicKeys: []provider.GPGPublicKey{
				{
					KeyID:      "C1C252F5499702CB",
					ASCIIArmor: "... public key ...",
				},
			},
		},
	}

	act, err := json.Marshal(&r)
	require.NoError(t, err)

	exp := []byte(`{"protocols":["5.0"],"os":"linux","arch":"amd64","filename":"terraform-provider-example_linux_x86_64-1.1.8.zip","download_url":"https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_linux_x86_64-1.1.8.zip","shasums_url":"https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_1.1.8_SHA256SUMS.txt","shasums_signature_url":"https://repo.example.com/artifactory/linux-generic-local/terraform/providers/terraform-provider-example/terraform-provider-example_1.1.8_SHA256SUMS.txt.sig","shasum":"d7dddb0a94c4388e4e3bf5f68faea18c46eab8aaceaec8954b269a4a29f13c29","signing_keys":{"gpg_public_keys":[{"key_id":"C1C252F5499702CB","ascii_armor":"... public key ..."}]}}`)

	assert.Equal(t, exp, act)
}
