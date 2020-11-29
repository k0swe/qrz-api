[![PkgGoDev](https://pkg.go.dev/badge/github.com/k0swe/qrz-api)](https://pkg.go.dev/github.com/k0swe/qrz-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/k0swe/qrz-api)](https://goreportcard.com/report/github.com/k0swe/qrz-api)

# Go API client for QRZ.com

A GoLang client library for QRZ.com's XML subscription data service. The service 
provides real-time access to information from the QRZ.COM servers and databases.

This client library was generated based on the OpenAPI specification in
the `api/openapi.yaml` file. However, the API itself is not well-described by
OpenAPI, so the generated library is supplemented with `wrapper.go`.

A simple application to demonstrate how to integrate the library is located in
`cmd/qrz-lookup/main.go`.
 
A QRZ.com XML subscription is required to take full advantage of the API. A
description of subscription plans and rates is available on the 
[QRZ.com website](http://www.qrz.com/i/subscriptions.html).
