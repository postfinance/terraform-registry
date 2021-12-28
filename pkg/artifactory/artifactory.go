//go:generate go install golang.org/x/tools/cmd/goimports@latest
//go:generate go install github.com/postfinance/httpclient/cmd/httpclient-gen-go@latest
//go:generate httpclient-gen-go -path . -package artifactory -force

// Package artifactory implements a minimal artifactory client and the provider backend for the PostFinance Artifactory
package artifactory

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"path"

	"github.com/postfinance/httpclient"
)

const (
	// AQLPath is the API path for AQL requests
	AQLPath = "/api/search/aql"
)

// Artifact represents the artifact
type Artifact struct {
	Repo   string `json:"repo"`
	Path   string `json:"path"`
	Name   string `json:"name"`
	SHA256 string `json:"sha256"`
}

// QueryService interface
type QueryService interface {
	Items(context.Context, AQL) ([]Artifact, *http.Response, error)
}

// QueryImpl implements QueryService
type QueryImpl struct {
	client *httpclient.Client
}

var _ QueryService = &QueryImpl{}

// Items returns all items matching the AQL expression
func (s QueryImpl) Items(ctx context.Context, find AQL) ([]Artifact, *http.Response, error) {
	// the AQL query has to be text/plain but the answer will be application/json
	s.client.RequestCallback = func(r *http.Request) *http.Request {
		r.Header.Set("Content-Type", httpclient.ContentTypeText)
		r.Body = io.NopCloser(bytes.NewBuffer(find.Bytes()))

		return r
	}

	req, err := s.client.NewRequest(
		http.MethodPost,
		path.Join(s.client.BaseURL.Path, AQLPath),
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	type itemsFindResponse struct {
		Results []Artifact `json:"results"`
		Range   struct {
			StartPos int `json:"start_pos"`
			EndPos   int `json:"end_pos"`
			Total    int `json:"total"`
		} `json:"range"`
	}

	res := &itemsFindResponse{}

	resp, err := s.client.Do(ctx, req, res)
	if err != nil {
		return nil, resp, err
	}

	// check EndPos/Total to ensure everything was loaded

	return res.Results, resp, nil
}
