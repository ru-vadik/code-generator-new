.PHONY: run
run:
	go run ./cmd/cg

.PHONY: build
build:
	go build -v ./cmd/cg

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test -v -timeout 60s ./...

.DEFAULT_GOAL := run
