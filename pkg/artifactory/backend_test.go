package artifactory

import (
	"net/url"
	"testing"

	"github.com/postfinance/terraform-registry/pkg/provider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProcessZIP(t *testing.T) {
	baseURL := "https://repo.example.com/artifactory"
	repo := "generic"
	repoPath := "terraform/providers"
	name := "terraform-provider-example_linux_amd64-0.0.1.zip"
	pType := "example"

	u, err := url.Parse(baseURL)
	require.NoError(t, err)

	s := Providers{
		url: u,
	}

	a := Artifact{
		Repo:   repo,
		Path:   repoPath,
		Name:   name,
		SHA256: "d7dddb0a94c4388e4e3bf5f68faea18c46eab8aaceaec8954b269a4a29f13c29",
	}

	t.Run("no error", func(t *testing.T) {
		exp := &provider.DownloadResponse{
			Protocols:           []string{APIVersion},
			OS:                  "linux",
			Arch:                "amd64",
			Filename:            name,
			DownloadURL:         baseURL + "/" + repo + "/" + repoPath + "/" + name,
			ShasumsURL:          baseURL + "/" + repo + "/" + repoPath + "/" + "terraform-provider-example_0.0.1_SHA256SUMS.txt",
			ShasumsSignatureURL: baseURL + "/" + repo + "/" + repoPath + "/" + "terraform-provider-example_0.0.1_SHA256SUMS.txt.sig",
			Shasum:              "d7dddb0a94c4388e4e3bf5f68faea18c46eab8aaceaec8954b269a4a29f13c29",
			SigningKeys: provider.SigningKeys{
				GPGPublicKeys: s.publicKeys,
			},
		}

		version, act, err := s.processZIP(a, pType)
		require.NoError(t, err)

		assert.Equal(t, "0.0.1", version)
		assert.Equal(t, exp, act)
	})

	t.Run("schema error 1", func(t *testing.T) {
		_, _, err := s.processZIP(Artifact{
			Name: "terraform-provider-example_linux_x86_64_1.1.8.zip",
		}, pType)

		assert.Error(t, err)
	})

	t.Run("schema error 2", func(t *testing.T) {
		_, _, err := s.processZIP(Artifact{
			Name: "terraform-provider-example_linux-1.1.8.zip",
		}, pType)

		assert.Error(t, err)
	})
}

func TestBuildURL(t *testing.T) {
	baseURL := "https://repo.example.com/artifactory"
	repo := "generic"
	repoPath := "terraform/providers"
	name := "terraform-provider-example_linux_amd64-0.0.1.zip"

	u, err := url.Parse(baseURL)
	require.NoError(t, err)

	s := Providers{
		url: u,
	}

	a := Artifact{
		Repo: repo,
		Path: repoPath,
	}

	act := s.buildURL(a, name)
	exp := baseURL + "/" + repo + "/" + repoPath + "/" + name

	assert.Equal(t, exp, act)
}
