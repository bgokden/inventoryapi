#!/usr/bin/env bash

go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen
mkdir -p api
# oapi-codegen  --package api --generate types spec/api.yaml  > api/api.gen.go

oapi-codegen  -package api -generate types -o api/types.gen.go spec/api.yaml  
oapi-codegen  -package api -generate server -o api/server.gen.go spec/api.yaml
oapi-codegen  -package api -generate spec -o api/spec.gen.go spec/api.yaml
oapi-codegen  -package api -generate client -o api/client.gen.go spec/api.yaml 
