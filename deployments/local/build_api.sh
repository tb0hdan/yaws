#!/usr/bin/env sh

# Build the API
oapi-codegen -config /app/api/config.yaml /app/api/api.yaml
cp /app/api/api.yaml /app/docs/api.yaml
chmod 666 /app/internal/server/api/api.gen.go
