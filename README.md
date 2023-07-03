![Xendit PHP SDK](docs/header.jpg "Xendit Go SDK")

# Xendit Go SDK

The official Xendit Go SDK provides a simple and convenient way to call Xendit's REST API
in applications written in Go.

* Package version: 2.0.0

# Getting Started

## Installation

Install xendit-go with:

```shell
go get github.com/kennycyb/xendit-go
```

Put the package under your project folder and add the following in import:

```golang
import xendit "github.com/kennycyb/xendit-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Authorization

The SDK needs to be instantiated using your secret API key obtained from the [Xendit Dashboard](https://dashboard.xendit.co/settings/developers#api-keys).
You can sign up for a free Dashboard account [here](https://dashboard.xendit.co/register).

# Documentation

Find detailed API infomration and examples for each of our product's by clicking the links below,

* [Customer](docs/CustomersApi.md)
* [Balance](docs/BalancesApi.md)

All URIs are relative to *https://api.xendit.co*.  For more information about our API, please refer to *https://developers.xendit.co/*.

Further Reading

* [Xendit Docs](https://docs.xendit.co/)
* [Xendit API Reference](https://developers.xendit.co/)