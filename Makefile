.PHONY: run
run:
	go run ./cmd/cg

.PHONY: run_race
run_race:
	go run -race ./cmd/cg

.PHONY: build
build:
	go build -v ./cmd/cg

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: test
test:
	go test -v -timeout 180s ./...

.DEFAULT_GOAL := run
