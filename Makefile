APP_NAME  := gonautobot
REPO      := github.com/josh-silvas
VERSION   := v$(shell cat VERSION)
HASH      := $(shell git rev-parse HEAD)
TS        := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')


# Small function to run golint for the packages directories
.PHONY: lint
lint: ## Run golint on all sub-packages
	@echo "Running linters on all sub-packages\n"
	golangci-lint run

.PHONY: release
release: ## Release a new version and build, tag.
	golangci-lint run
	rm -rf build/*
	git tag $(VERSION)
	goreleaser --rm-dist

.PHONY: testrelease
testrelease: ## Test a release
	git tag $(version)
	goreleaser --rm-dist --snapshot
	git tag -d $(version)

.PHONY: test
test:
	@go test -v -short -coverprofile coverage.txt -covermode atomic $(REPO)/$(APP_NAME)/... | { grep -v 'no test files'; true; }

.PHONY: cover
cover:
	@go tool cover -html=coverage.txt


