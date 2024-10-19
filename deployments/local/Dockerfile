FROM golang:1.23-alpine AS builder
WORKDIR /
RUN apk update && apk add --no-cache go
RUN go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
