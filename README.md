# CDKTF Go bindings for hashicorp/time provider version 0.13.1

This repo builds and publishes the [Terraform time provider](https://registry.terraform.io/providers/hashicorp/time/0.13.1/docs) bindings for [CDK for Terraform](https://cdk.tf).

## Go Package

The go package is generated into the [`github.com/cdktn-io/cdktn-provider-time-go`](https://github.com/cdktn-io/cdktn-provider-time-go) package.

`go get github.com/cdktn-io/cdktn-provider-time-go/time/<version>`

Where `<version>` is the version of the prebuilt provider you would like to use e.g. `v11`. The full module name can be found
within the [go.mod](https://github.com/cdktn-io/cdktn-provider-time-go/blob/main/time/go.mod#L1) file.

## Docs

Find auto-generated docs for this provider [here](https://github.com/cdktn-io/cdktn-provider-time/blob/main/docs/API.go.md).


## Versioning

This project is explicitly not tracking the Terraform time provider version 1:1. In fact, it always tracks `latest` of `~> 0.7` with every release. If there are scenarios where you explicitly have to pin your provider version, you can do so by [generating the provider constructs manually](https://cdk.tf/imports).

These are the upstream dependencies:

* [CDK for Terraform](https://cdk.tf) - Last official release
* [Terraform time provider](https://registry.terraform.io/providers/hashicorp/time/0.13.1)
* [Terraform Engine](https://terraform.io)

If there are breaking changes (backward incompatible) in any of the above, the major version of this project will be bumped.

## Features / Issues / Bugs

Please report bugs and issues to the [CDK for Terraform](https://cdk.tf) project:

* [Create bug report](https://cdk.tf/bug)
* [Create feature request](https://cdk.tf/feature)

## Contributing

### Projen

This is mostly based on [Projen](https://projen.io), which takes care of generating the entire repository.

### cdktn-provider-project based on Projen

There's a custom [project builder](https://github.com/cdktn-io/cdktn-provider-project) which encapsulate the common settings for all `cdktf` prebuilt providers.


### Repository Management

The repository is managed by [CDKTN Repository Manager](https://github.com/cdktn-io/cdktn-repository-manager/).
