.PHONY: build

OUTPUT = {{cookiecutter.project_name}}
VERSION = $(shell git describe --tags --always --dirty)

build:
	go build -i -v -o build/${OUTPUT} -ldflags="-X main.version=${VERSION}"

build-linux:
	GOOS=linux GOARCH=amd64 go build -i -v -o build/${OUTPUT}-linux-amd64 -ldflags="-X main.version=${VERSION}"

install:
	go install -i -v -ldflags="-X main.version=${VERSION}"
