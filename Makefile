VERSION := $(shell git describe --tags --abbrev=0)
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

docker-build:
	@docker build --tag jfosse/fizzbuzzhttp:latest --tag jfosse/fizzbuzzhttp:${VERSION} --force-rm --build-arg VERSION=${VERSION} .

integration-test: docker-build
	@docker run --rm --name fizzbuzzhttpIT -d -p 8080:8080 jfosse/fizzbuzzhttp:${VERSION}
	@go test -v ./itest/fizzbuzz_test.go -tags=integration
	@docker stop fizzbuzzhttpIT
	@docker run --rm --name fizzbuzzhttpIT -d -p 8080:8080 jfosse/fizzbuzzhttp:${VERSION}
	@go test -v ./itest/statistics_test.go -tags=integration
	@docker stop fizzbuzzhttpIT

docker-push: integration-test
	@docker push jfosse/fizzbuzzhttp:latest
	@docker push jfosse/fizzbuzzhttp:${VERSION}	
