[![Actions Status](https://github.com/pulumi/pulumi-http/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-http/actions)
[![NPM version](https://img.shields.io/npm/v/@pulumi/http)](https://www.npmjs.com/package/@pulumi/http)
[![Python version](https://img.shields.io/pypi/v/pulumi_http)](https://pypi.org/project/pulumi_http)
[![NuGet version](https://img.shields.io/nuget/v/Pulumi.Http)](https://www.nuget.org/packages/Pulumi.Http)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-http/sdk/go)](https://pkg.go.dev/github.com/pulumi/pulumi-http/sdk/go)
[![License](https://img.shields.io/github/license/pulumi/pulumi-http)](https://github.com/pulumi/pulumi-http/blob/master/LICENSE)

# HTTP Resource Provider

This provider is mainly used for ease of converting terraform programs to Pulumi.
For standard use in Pulumi programs, please use your programming language's standard http library.

The HTTP resource provider for Pulumi lets you use HTTP resources in your cloud programs.
To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/install/).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/http

or `yarn`:

    $ yarn add @pulumi/http

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_http

### Go

To use from Go, use `go get` to grab the latest version of the library:

    $ go get github.com/pulumi/pulumi-http/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Http

<!-- If your provider has configuration, remove this comment and the comment tags below, updating the documentation. -->
<!--

## Configuration

The following Pulumi configuration can be used:

- `http:token` - (Required) The API token to use with HTTP. When not set, the provider will use the `HTTP_TOKEN` environment variable.

-->

<!-- If your provider has reference material available elsewhere, remove this comment and the comment tags below, updating the documentation. -->
<!--

## Reference

For further information, please visit [HTTP reference documentation](https://example.com/http).

-->
