export

.PHONY: all build up

GOPATH ?= $(HOME)/go

up:
	docker-compose up -d

build:
	go build -o user-generator -a -v -tags musl

mod:
	go mod download

api:
	go run -tags musl -race main.go

down:
	docker-compose down --remove-orphans

clean:
	go mod tidy

test:
	go test -v -count=1 ./...

test-cover:
	go test -v -count=1 ./... -covermode=count

test-out:
	go test -v -coverprofile cover.out ./...
	go tool cover -html=cover.out -o cover.html

run-api:
	$(GOPATH)/bin/reflex -s -r '\.go$$' make api

run: run-api