VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

default: build

clean: 
	@echo "Cleaning..."
	@go clean

test:
	@echo "Testing..."
	@go test -v ./...

check-version:
ifndef ${VERSION}
VERSION = v0.0.0
endif

build: clean test check-version
	@echo "Building..."
	@go build -i -o ${PROJECTNAME} -v -ldflags "-X main.version=${VERSION} -X main.build=${BUILD}"
	@echo "Done."

install: build
	@echo "Installing "${PROJECTNAME}" in "${GOBIN}
	@go install	
