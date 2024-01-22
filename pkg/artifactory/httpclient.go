// Code generated by client-gen-go; DO NOT EDIT.
// This file was generated by robots at
// 2024-01-22 11:27:51.246200827 +0100 CET m=+0.001714141

package artifactory

import "github.com/postfinance/httpclient"

// Client is a generated wrapper for a http client and detected services.
type Client struct {
	*httpclient.Client

	// Services used for communicating with the API
	Query QueryService
}

// NewClient returns a new API client.
func NewClient(baseURL string, opts ...httpclient.Opt) (*Client, error) {

	client, err := httpclient.New(baseURL, opts...)
	if err != nil {
		return nil, err
	}

	// services
	query := &QueryImpl{client: client}

	return &Client{
		client,
		query,
	}, nil
}
