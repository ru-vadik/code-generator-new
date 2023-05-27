.PHONY: run
run:
	go run ./cmd/cg

.PHONY: build
build:
	go build -v ./cmd/cg

.PHONY: vendor
vendor:
	go mod vendor

.DEFAULT_GOAL := run
