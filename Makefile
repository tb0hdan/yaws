.PHONY: api

DEV_IMAGE=yaws-dev-image

.image:
	@docker build -t $(DEV_IMAGE) -f deployments/local/Dockerfile .
	@touch .image

api: .image
	@docker run --rm -it -v $(PWD):/app -w /app $(DEV_IMAGE) /bin/sh /app/deployments/local/build_api.sh

lint:
	@golangci-lint run ./...

run:
	@go run ./cmd/app/main.go
