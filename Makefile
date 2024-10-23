# Variables
BINARY := simple-env-parser
DOCKER_REPO := sagoresarker/simple-env-parser
GOBIN := $(PWD)/bin

GO := go
DOCKER := docker

$(shell mkdir -p $(GOBIN))

.PHONY: all build test clean docker-*

all: test build

build:
	@ $(GO) build -o $(GOBIN)/$(BINARY)

test:
	@ $(GO) test -v ./...

clean:
	@ $(GO) clean
	@ rm -rf $(GOBIN)

docker-build:
	@ $(DOCKER) build -t $(DOCKER_REPO):local .

docker-build-ci:
	@ $(DOCKER) build -t $(DOCKER_REPO):$(tag) .
	@ $(DOCKER) build -t $(DOCKER_REPO):latest .

docker-push-ci:
	@ $(DOCKER) push $(DOCKER_REPO):$(tag)
	@ $(DOCKER) push $(DOCKER_REPO):latest

help:
	@ echo "Available targets:"
	@ echo "  build         - Build binary"
	@ echo "  test         - Run tests"
	@ echo "  clean        - Remove built files"
	@ echo "  docker-build - Build local Docker image"
	@ echo "  docker-*-ci  - CI Docker operations"