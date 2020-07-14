#!/bin/bash
openapi-generator generate -i=api/openapi.yaml -g=go -c=api/openapi-config.yaml
go fmt ./...
go mod tidy
