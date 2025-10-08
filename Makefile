.PHONY: run-api run-worker migrate-up migrate-down lint docker-build docker-run-api docker-run-worker

run-api:
	go run ./cmd/api

run-worker:
	go run ./cmd/worker

migrate-up:
	goose -dir ./migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir ./migrations postgres "$(DATABASE_URL)" down

lint:
	gofmt -s -w . && go vet ./...

docker-build:
	docker build -t stock-screener:local .

docker-run-api:
	docker run --env-file .env -p 8080:8080 stock-screener:local api

docker-run-worker:
	docker run --env-file .env stock-screener:local worker
