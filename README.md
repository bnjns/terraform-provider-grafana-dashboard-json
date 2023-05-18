<div align="center">

### Grafana Dashboard JSON Terraform Provider

![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/bnjns/terraform-provider-grafana-dashboard-json/build.yml?branch=main&style=flat-square)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/bnjns/terraform-provider-grafana-dashboard-json?display_name=tag&label=version&sort=semver&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/bnjns/terraform-provider-grafana-dashboard-json?style=flat-square)

---

A Terraform provider that lets you manage your Metabase instance, because why not Terraform the world?
</div>

## 🧐 About

<!-- TODO -->

## 🏁 Getting Started

### Prerequisites

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.19

### Installing

Simply clone this repo to your desired location:

```sh
$ git clone git@github.com:bnjns/terraform-provider-grafana-dashboard-json.git
```

Install the Go dependencies:

```sh
$ go mod download
```

## 🎈 Usage

### Building the provider

To build the provider and install into your `GOPATH`:

```sh
$ go install
```

### Configuring Terraform

You can configure Terraform to use a [local build](#building-the-provider) by adding the following to
you `~/.terraformrc` file:

```hcl
provider_installation {
  dev_overrides {
    "bnjns/grafana-dashboard-json" = "</path/to/GOPATH>/bin"
  }

  direct {}
}
```

> **Note:** You must include `direct {}` otherwise all other providers will fail to install.

### Generating the documentation

The documentation can be auto-generated using `tfplugindocs`:

```sh
go generate
```

### Running the tests

To run the unit tests:

```sh
$ go test -v ./...
```

To run the provider acceptance tests:

```sh
$ TF_ACC=1 go test -v ./... -run "^TestAcc"
```

## 🚀 Releasing

Releasing is handled automatically by [GitHub Actions](.github/workflows/release.yml) and
Hashicorp's `terraform-provider-release` action. An admin will simply tag the latest release to trigger the pipeline.

## ⛏️ Built Using

- [terraform-provider-scaffolding-framework](https://github.com/hashicorp/terraform-provider-scaffolding-framework)

## ✍️ Authors

- [@bnjns](https://github.com/bnjns)
