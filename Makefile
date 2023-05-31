.PHONY: run
run:
	go run ./cmd/cg

.PHONY: run_race
run_race:
	go run -race ./cmd/cg

.PHONY: run_prof
run_prof:
	go run ./cmd/cg -cpuProfile

.PHONY: prof_web
prof_web:
	go tool pprof -http=localhost:8080 cpuProfile.prof

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
