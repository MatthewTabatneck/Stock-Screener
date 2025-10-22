include .env
export

.PHONY: run-api run-worker migrate-up migrate-down lint docker-build docker-run-api docker-run-worker

run-api: #locally run api
	go run ./cmd/api

run-worker: #locally run worker
	go run ./cmd/worker

migrate-up: #Initialize db or merge new changes
	goose -dir ./migrations postgres "$(DATABASE_URL)" up

migrate-down: #Remove tables nuke db
	goose -dir ./migrations postgres "$(DATABASE_URL)" down

lint: ##clean code
	gofmt -s -w . && go vet ./...

docker-build:
	docker build -t stock-screener:local .

docker-run-api:
	docker run --env-file .env -p 8080:8080 stock-screener:local api

docker-run-worker:
	docker run --env-file .env stock-screener:local worker

up:
	docker compose up --build

down:  
	docker compose down

logs:  
	docker compose logs -f
