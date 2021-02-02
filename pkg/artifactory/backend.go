package artifactory

import (
	"context"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/marcsauter/terraform-registry/pkg/provider"
	"github.com/postfinance/httpclient"
)

// Constants
const (
	Namespace = "postfinance"
	Repo      = "linux-generic-local"
	RepoPath  = "terraform/providers"
)

// Providers implements provider.Backend for Artifactory
type Providers struct {
	url    *url.URL
	client *Client
}

// New return a new provider.Backend for Artifactory
func New(baseURL, username, password string) (*Providers, error) {
	c, err := NewClient(baseURL,
		httpclient.WithUsername(username),
		httpclient.WithPassword(password),
	)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// modify path to point to the location of the terraform providers
	u.Path = path.Join(u.Path)

	return &Providers{
		url:    u,
		client: c,
	}, nil
}

var _ provider.Backend = &Providers{}

// Versions implements provider.Backend
func (s Providers) Versions(req *provider.VersionsRequest) (*provider.VersionsResponse, error) {
	if req.Namespace != Namespace {
		return nil, fmt.Errorf("namespace %q unknown", req.Namespace)
	}

	// get all artifacts of the requested type
	artifacts, err := s.getProviders(req.Type)
	if err != nil {
		return nil, err
	}

	versions := make(map[string]*provider.Version)
	// consolidate all platforms per version
	for version, providers := range artifacts {
		v, ok := versions[version]
		if !ok {
			v = &provider.Version{
				Version:   version,
				Protocols: []string{},
				Platforms: []provider.Platform{},
			}
			versions[version] = v
		}

		// get the platform of all providers
		for _, p := range providers {
			v.Platforms = append(v.Platforms, provider.Platform{
				OS:   p.OS,
				Arch: p.Arch,
			})
		}
	}

	return &provider.VersionsResponse{
		Versions: func() []provider.Version {
			res := []provider.Version{}

			for _, v := range versions {
				res = append(res, *v)
			}

			return res
		}(),
	}, nil
}

// Download implements provider.Backend
// curl 'https://registry.terraform.io/v1/providers/hashicorp/random/2.0.0/download/linux/amd64'
func (s Providers) Download(req *provider.DownloadRequest) (*provider.DownloadResponse, error) {
	if req.Namespace != Namespace {
		return nil, fmt.Errorf("namespace %q unknown", req.Namespace)
	}

	// get all artifacts of the requested type
	providers, err := s.getProviders(req.Type)
	if err != nil {
		return nil, err
	}

	// return the provider of the requested version, os and arch
	for _, r := range providers[req.Version] {
		if r.OS == req.OS && r.Arch == req.Arch {
			return r, nil
		}
	}

	return nil, fmt.Errorf("requested provider not found")
}

// getProviders of the type t from Artifactory
func (s Providers) getProviders(t string) (map[string][]*provider.DownloadResponse, error) {
	find := FindItems(Repo, path.Join(RepoPath, fmt.Sprintf("terraform-provider-%s", t)), "*")

	artifacts, resp, err := s.client.Query.Items(context.TODO(), find)
	if err != nil {
		return nil, err
	}

	_ = resp.Body.Close()

	providers := make(map[string][]*provider.DownloadResponse)

	for _, a := range artifacts {
		if path.Ext(a.Name) != ".zip" {
			continue
		}

		v, p, err := s.processZIP(a, t)
		if err != nil {
			return nil, err
		}

		providers[v] = append(providers[v], p)
	}

	return providers, nil
}

func (s Providers) processZIP(a Artifact, ptype string) (string, *provider.DownloadResponse, error) {
	var schemaError = fmt.Errorf("name of .zip file does not match schema terraform-provider-:type:_:os:_:arch:-:version:.zip: %s", a.Name)

	// "terraform-provider-aixboms_linux_x86_64-1.1.8.zip"
	name := strings.TrimSuffix(a.Name, ".zip")
	// "terraform-provider-aixboms_linux_x86_64-1.1.8"
	part := strings.Split(name, "-")
	// []string{"terraform", "provider", "aixboms_linux_x86_64", "1.1.8"}

	if len(part) != 4 {
		return "", nil, schemaError
	}

	version := part[len(part)-1]                                              // "1.1.8"
	p := strings.SplitN(strings.TrimPrefix(part[len(part)-2], ptype), "_", 3) // []string{"", "linux", "x86_64"}

	if len(p) != 3 {
		return "", nil, schemaError
	}

	const (
		sha256sums = `terraform-provider-%s_%s_SHA256SUMS.txt%s`
	)

	res := &provider.DownloadResponse{
		Protocols:           []string{},
		OS:                  p[1],
		Arch:                replaceArch(p[2]),
		Filename:            a.Name,
		DownloadURL:         s.buildURL(a, a.Name),
		ShasumsURL:          s.buildURL(a, fmt.Sprintf(sha256sums, ptype, version, "")),     // terraform-provider-aixboms_1.1.8_SHA256SUMS.txt
		ShasumsSignatureURL: s.buildURL(a, fmt.Sprintf(sha256sums, ptype, version, ".sig")), // terraform-provider-aixboms_1.1.8_SHA256SUMS.txt.sig
		Shasum:              a.SHA256,
		SigningKeys:         provider.SigningKeys{},
	}

	return version, res, nil
}

func (s Providers) buildURL(a Artifact, n string) string {
	return fmt.Sprintf("%v/%s", s.url, path.Join(a.Repo, a.Path, n))
}

func replaceArch(p string) string {
	const (
		amd64  = "amd64"
		x86_64 = "x86_64"
	)

	switch p {
	case x86_64:
		return amd64
	case amd64:
		return x86_64
	default:
		return p
	}
}
