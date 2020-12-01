DIRS := $(shell go list ./...)

.PHONY: help deps fmt lint test test-race test-integration

help:
	@echo ""
	@echo "Welcome to DataDog/jamf-api-client-go make."
	@echo "The following commands are available:"
	@echo ""
	@echo "    make clean             : Cleanup all test files and binaries"
	@echo "    make deps              : Fetch all dependencies"
	@echo "    make fmt               : Run go fmt to fix any formatting issues"
	@echo "    make lint              : Use go vet to check for linting issues"
	@echo "    make test              : Run all short tests"
	@echo "    make test-race         : Run all tests with race condition checking"
	@echo "    make test-integration  : Run all tests without limiting to short"
	@echo ""
	@echo "    make pr-prep           : Run this before making a PR to run fmt, lint and tests"
	@echo ""

clean:
	rm -f cp.out
	rm -f .coverage.html
	rm -rf bin/
	rm -rf vendor/

deps:
	@go mod tidy

fmt:
	@go fmt ${DIRS}

lint:
	@go vet ${DIRS}

test:
	@go test -v -coverprofile=cp.out  -count=1 -timeout 300s -short ${DIRS}
	go tool cover -html=cp.out -o .coverage.html

test-race:
	@go test -v -coverprofile=cp.out  -count=1 -timeout 300s -short -race ${DIRS}
	go tool cover -html=cp.out -o .coverage.html

test-integration:
	@go test -v -coverprofile=cp.out  -count=1 -timeout 600s ${DIRS}
	go tool cover -html=cp.out -o .coverage.html

pr-prep: clean fmt lint test-race test-integration
