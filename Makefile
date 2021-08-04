app_name := josh5276/gonautobot
version := v$(shell cat VERSION)
hash := $(shell git rev-parse HEAD)
timestamp := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')


# Small function to run golint for the packages directories
.PHONY: lint
lint: ## Run golint on all sub-packages
	@echo "Running linters on all sub-packages\n"
	golangci-lint run

.PHONY: release
release: ## Release a new version and build, tag.
	golangci-lint run
	rm -rf build/*
	git tag $(version)
	goreleaser --rm-dist

.PHONY: testrelease
testrelease: ## Test a release
	git tag $(version)
	goreleaser --rm-dist --snapshot
	git tag -d $(version)

.PHONY: test
test:
	go test -v -short github.com/$(app_name)/common


