GO            ?= go
COVERAGE_FILE := coverage.out

.DEFAULT_GOAL := help

.PHONY: help test cover vet lint fmt tidy check

help:
	@echo "Available targets:"
	@echo "  make test                          Run tests"
	@echo "  make cover                         Run tests with coverage"
	@echo "  make vet                           Run go vet"
	@echo "  make lint                          Run golangci-lint"
	@echo "  make fmt                           Format Go source files"
	@echo "  make tidy                          Synchronize module dependencies"
	@echo "  make check                         Run test, vet, and lint"

test:
	$(GO) test ./...

cover:
	$(GO) test -coverprofile=$(COVERAGE_FILE) ./...
	$(GO) tool cover -func=$(COVERAGE_FILE)

vet:
	$(GO) vet ./...

lint:
	golangci-lint run

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

check: test vet lint
