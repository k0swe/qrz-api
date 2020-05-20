#!/bin/bash
openapi-generator generate -i=api/openapi.yaml -g=go -c=api/openapi-config.yaml
