VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $(shell basename "$(PWD)")

ifndef VERSION
	VERSION = 0.0.0
endif

default: build

clean: 
	@echo "Cleaning..."
	@go clean

test:
	@echo "Testing..."
	@go test -v ./...

build: clean test
	@echo "Building version ${VERSION}"
	@go build -i -o ${PROJECTNAME} -v -ldflags "-X main.version=${VERSION} -X main.build=${BUILD}"
	@echo "Done."

install: build
	@echo "Installing "${PROJECTNAME}" in "${GOBIN}
	@go install	

build-image:
	@docker build --tag jfosse/fizzbuzzhttp:latest --tag jfosse/fizzbuzzhttp:${VERSION} .

push-image: build-image
	@docker push jfosse/fizzbuzzhttp:latest	
	@docker push jfosse/fizzbuzzhttp:${VERSION}	