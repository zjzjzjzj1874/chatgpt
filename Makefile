PKG = $(shell cat go.mod | grep "^module " | sed -e "s/module //g")
VERSION = v$(shell cat cmd/version)
COMMIT_SHA ?= $(shell git describe --always)-devel

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
GOBUILD=CGO_ENABLED=0 go build -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"
GOINSTALL=CGO_ENABLED=0 go install -ldflags "-X ${PKG}/version.Version=${VERSION}+sha.${COMMIT_SHA}"

MAIN_ROOT ?= ./cmd

install: download
	cd $(MAIN_ROOT) && $(GOINSTALL)

build:
	cd $(MAIN_ROOT) && $(GOBUILD) -o gptx

download:
	go mod download -x

clean:
	rm -rf ./cmd/$(WORKSPACE)/out

upgrade:
	go get -u ./...

tidy:
	go mod tidy

docker:
	docker buildx build --push --progress plain --platform=${PLATFORM}	\
		--cache-from "type=local,src=/tmp/.buildx-cache" \
		--cache-to "type=local,dest=/tmp/.buildx-cache" \
		--file=./Dockerfile \
		--tag=gptx:${VERSION} \
		.