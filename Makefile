.PHONY:

APPLICATION_NAME ?= srcm

run: build
	@./bin/${APPLICATION_NAME}

build:
	@go build -o bin/${APPLICATION_NAME}

test:
	@go test -v ./...
	