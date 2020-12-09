//go:generate httpclient-gen-go -path . -package artifactory -force
// Package artifactory implements a minimal artifactory client
package artifactory

import (
	"context"
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
	Repo string `json:"repo"`
	Path string `json:"path"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size int    `json:"size"`
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
	req, err := s.client.NewRequest(
		http.MethodPost,
		path.Join(s.client.BaseURL.String(), AQLPath),
		find.String(),
	)
	if err != nil {
		return nil, nil, err
	}

	// has to be text/plain for AQL
	req.Header.Add("Content-Type", httpclient.ContentTypeText)

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
