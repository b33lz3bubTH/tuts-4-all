all: build

build:
	go build -o bin/tuts4all ./cmd/main.go

run:
	go run ./cmd/main.go

migrate:
	go run ./cmd/main.go migrate

docker-build:
	docker build -t tuts4all-backend .

docker-run:
	docker run --env-file=config.env -p 8080:8080 tuts4all-backend

.PHONY: run migrate
