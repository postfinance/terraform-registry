package artifactory_test

import (
	"context"
	"os"
	"testing"

	"github.com/postfinance/httpclient"
	"github.com/postfinance/terraform-registry/pkg/artifactory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestItems(t *testing.T) {
	url, ok := os.LookupEnv("ARTIFACTORY_BASE_URL")
	if !ok {
		t.Skip("environment variable ARTIFACTORY_BASE_URL is not set")
	}
	u, ok := os.LookupEnv("ARTIFACTORY_USERNAME")
	if !ok {
		t.Skip("environment variable ARTIFACTORY_USERNAME is not set")
	}
	p, ok := os.LookupEnv("ARTIFACTORY_PASSWORD")
	if !ok {
		t.Skip("environment variable ARTIFACTORY_PASSWORD is not set")
	}

	c, err := artifactory.NewClient(url,
		httpclient.WithUsername(u),
		httpclient.WithPassword(p),
	)
	require.NoError(t, err)
	require.NotNil(t, c)

	find := artifactory.FindItems("linux-generic-local", "terraform/providers", "terraform-provider-uam*.zip")
	a, resp, err := c.Query.Items(context.Background(), find)
	assert.NotEmpty(t, a)
	assert.NotNil(t, resp)
	assert.NoError(t, err)
}
