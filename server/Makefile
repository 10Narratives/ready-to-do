BINARY_NAME=server

GO_FILES=$(shell find . -type f -name '*.go')
PKG_LIST=$(shell go list ./...)

.PHONY: all build run clean lint tests

all: build

build:
	@GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux ./...
	@GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows.exe ./...
	@GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin ./...

run:
	@go run main.go

clean:
	@go clean
	@rm -f ${BINARY_NAME}
	@rm -f coverage.out

tests:
	@go test -v -coverprofile=coverage.out ${PKG_LIST}
	@go tool cover -func=coverage.out

deps:
	@go mod download