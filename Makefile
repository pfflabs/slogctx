MAKEFILE_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
BIN_DIR := $(MAKEFILE_DIR)/bin
GOLINT_VERSION := v1.54.2
GOLINT := $(BIN_DIR)/golangci-lint-$(GOLINT_VERSION)
GINKGO_VERSION := $(shell cat go.mod | grep github.com/onsi/ginkgo/v2 | cut -d " " -f 2)
ifeq ($(GINKGO_VERSION),)
	$(error did not find ginkgo version in go.mod)
endif
GINKGO := $(BIN_DIR)/ginkgo-$(GINKGO_VERSION)
COVERAGE_OUT := $(MAKEFILE_DIR)/coverprofile.out

.PHONY: check
check: vet lint test

.PHONY: coverage
coverage: $(COVERAGE_OUT)

.PHONY: view-coverage
view-coverage: $(COVERAGE_OUT)
	@go tool cover -html="$(COVERAGE_OUT)"

.PHONY: test
test: $(GINKGO)
	@$(GINKGO) run $(GINKGO_OPTS) "$(MAKEFILE_DIR)/..."

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint: $(GOLINT)
	@$(GOLINT) run

$(COVERAGE_OUT): override GINKGO_OPTS += --cover
$(COVERAGE_OUT): test

$(GOLINT):
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s $(GOLINT_VERSION)
	@mv "$(BIN_DIR)/golangci-lint" "$(GOLINT)"

$(GINKGO):
	@GOBIN="$(BIN_DIR)" go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)
	@mv "$(BIN_DIR)/ginkgo" "$(GINKGO)"
