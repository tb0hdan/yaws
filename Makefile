.PHONY: api mocks

DEV_IMAGE=yaws-dev-image

.image:
	@docker build -t $(DEV_IMAGE) -f deployments/local/Dockerfile .
	@touch .image

api: .image
	@docker run --rm -it -v $(PWD):/app -w /app $(DEV_IMAGE) /bin/sh /app/deployments/local/build_api.sh

mocks: .image
	@docker run --rm -it -v $(PWD):/app -w /app $(DEV_IMAGE) /bin/sh /app/deployments/local/build_mocks.sh

lint:
	@golangci-lint run ./...

run:
	@go run ./cmd/app/main.go

dbup:
	@docker compose -f deployments/local/docker-compose.yaml  up -d --remove-orphans

dbdown:
	@docker compose -f deployments/local/docker-compose.yaml down

dbpurge: dbdown
	@docker volume rm local_yaws_db_data local_yaws_db_backups
