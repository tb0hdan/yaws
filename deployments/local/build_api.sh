#!/usr/bin/env sh

if [ -z "$1" ]; then
    echo "User/Group required"
    exit 127
fi

# Build the API
oapi-codegen -config /app/api/config.yaml /app/api/api.yaml
cp /app/api/api.yaml /app/docs/api.yaml
chown -R $1 /app/internal/server/api/api.gen.go
chown -R $1 /app/docs
