[![Release](https://img.shields.io/github/release/postfinance/terraform-registry.svg?style=for-the-badge)](https://github.com/postfinance/terraform-registry/releases/latest)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=for-the-badge)](/LICENSE)
[![Build](https://img.shields.io/github/workflow/status/postfinance/terraform-registry/build?style=for-the-badge)](https://github.com/postfinance/terraform-registry/actions?query=workflow%3Abuild)
[![Go Report Card](https://img.shields.io/badge/GOREPORT-A%2B-brightgreen.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/postfinance/terraform-registry)

# terraform-registry

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [terraform-registry](#terraform-registry)
  - [Definitions](#definitions)
  - [Provider Registry](#provider-registry)
    - [Protocol Versions](#protocol-versions)
    - [Acceptance Testing](#acceptance-testing)
  - [Tests](#tests)
    - [Service discovery](#service-discovery)
    - [Provider versions](#provider-versions)
    - [Provider download](#provider-download)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->



## Definitions
- [Module Registry Protocol](https://www.terraform.io/docs/internals/module-registry-protocol.html)
- [Provider Registry Protocol](https://www.terraform.io/docs/internals/provider-registry-protocol.html)

## Provider Registry
### Protocol Versions 

This is about the Terraform provider API versions: 

- https://www.terraform.io/docs/internals/provider-registry-protocol.html#protocols
- https://www.terraform.io/docs/internals/provider-registry-protocol.html#protocols-1

> [see this Reddit post for details](https://www.reddit.com/r/Terraform/comments/iydnpq/figuring_out_protocol_version/g6metb6/?utm_source=share&utm_medium=web2x&context=3)

### Acceptance Testing
For Artifactory acceptance testing set:
```shell
export ARTIFACTORY_BASE_URL=
export ARTIFACTORY_USERNAME=
export ARTIFACTORY_PASSWORD=
```

## Tests 

### Service discovery

```shell
curl -s http://localhost:8080/.well-known/terraform.json | jq
```
```json
{
  "providers.v1": "/v1/providers"
}
```

### Provider versions

```shell
curl -s ttp://localhost:8080/v1/providers/postfinance/example/versions | jq
```
```json
{
  "versions": [
    {
      "version": "0.0.1",
      "protocols": [
        "5.0"
      ],
      "platforms": [
        {
          "os": "linux",
          "arch": "amd64"
        }
      ]
    },
    {
      "version": "1.1.9",
      "protocols": [
        "5.0"
      ],
      "platforms": [
        {
          "os": "linux",
          "arch": "amd64"
        }
      ]
    }
  ]
}
```

### Provider download
```shell
curl -s http://localhost:8080/v1/providers/postfinance/example/0.0.1/download/linux/amd64 | jq
```
```json
{
  "protocols": [
    "5.0"
  ],
  "os": "linux",
  "arch": "amd64",
  "filename": "terraform-provider-example_linux_x86_64-0.0.1.zip",
  "download_url": "https://repo.example.com/artifactory/generic/terraform/providers/terraform-provider-example/terraform-provider-example_linux_x86_64-0.0.1.zip",
  "shasums_url": "https://repo.example.com/artifactory/generic/terraform/providers/terraform-provider-example/terraform-provider-example_0.0.1_SHA256SUMS.txt",
  "shasums_signature_url": "https://repo.example.com/artifactory/generic/terraform/providers/terraform-provider-example/terraform-provider-example_0.0.1_SHA256SUMS.txt.sig",
  "shasum": "d7dddb0a94c4388e4e3bf5f68faea18c46eab8aaceaec8954b269a4a29f13c29",
  "signing_keys": {
    "gpg_public_keys": [
      {
        "key_id": "C1C252F5499702CB",
        "ascii_armor": "-----BEGIN PGP PUBLIC KEY BLOCK-----\n ... -----END PGP PUBLIC KEY BLOCK-----\n"
      }
    ]
  }
}```
