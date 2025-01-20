#!/bin/bash
npm install @openapitools/openapi-generator-cli prettier -g
openapi-generator-cli generate -i=api/openapi.yaml -g=go -c=api/openapi-config.yaml
find . -type f -exec sed -i 's/GIT_USER_ID\/GIT_REPO_ID/k0swe\/qrz\-logbook/g' {} +
go get -u
go mod tidy
go fmt ./...
prettier -w --prose-wrap always .
