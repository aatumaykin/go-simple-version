GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
VERSION?="dev"

# The ldflags for the Go build process to set the version related data
GO_BUILD_VERSION_LDFLAGS=\
  -X github.com/aatumaykin/go-simple-version.Version=$(VERSION) \
  -X github.com/aatumaykin/go-simple-version.BuildTime=$(shell date +"%Y-%m-%dT%H:%M:%S%z") \
  -X github.com/aatumaykin/go-simple-version.Commit=$(shell git rev-parse --short HEAD)

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(GO_BUILD_LDFLAGS)" -o example ./example/
.PHONY: build