VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse HEAD)
DATE := $(shell date -u +%Y-%m-%d)

.PHONY: build
build:
	CGO_ENABLED=0 go build -ldflags="-s -w \
			-X main.version=${VERSION} \
			-X main.commit=${COMMIT} \
			-X main.date=${DATE}" \
			-o dist/ ./cmd/gofile

.PHONY: test
test:
	go test -v ./...
