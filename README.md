# Go API client for qrz

This document describes the interface specification for access to QRZ's XML subscription data service. This service provides real-time access to information from the QRZ.COM servers and databases.
Access to this service requires user authentication through the use of a valid QRZ.COM username and password. While any QRZ user may login to the service, an active QRZ Logbook Data subscription is required to access most of its features. Non-subscriber access limits the data fields that are returned and is primarily intended for testing and troubleshooting purposes only.
A description of subscription plans and rates is available on the [QRZ website](http://www.qrz.com/i/subscriptions.html).

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.34
- Package version: 0.0.3
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./qrz"
```

## Documentation for API Endpoints

All URIs are relative to *https://xmldata.qrz.com/xml/1.34*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**RootGet**](docs/DefaultApi.md#rootget) | **Get** / | The do-everything endpoint


## Documentation For Models

 - [Callsign](docs/Callsign.md)
 - [Dxcc](docs/Dxcc.md)
 - [QrzDatabase](docs/QrzDatabase.md)
 - [Session](docs/Session.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author

va7stv@qrz.com

