BINPATH ?= build

BUILD_TIME=$(shell date +%s)
GIT_COMMIT=$(shell git rev-parse HEAD)
VERSION ?= $(shell git tag --points-at HEAD | grep ^v | head -n 1)

LDFLAGS = -ldflags "-X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT) -X main.Version=$(VERSION)"

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

NVM_SOURCE_PATH ?= $(HOME)/.nvm/nvm.sh

ifneq ("$(wildcard $(NVM_SOURCE_PATH))","")
	NVM_EXEC = source $(NVM_SOURCE_PATH) && nvm exec --
endif
NPM = $(NVM_EXEC) npm

.PHONY: all
all: delimiter-AUDIT audit delimiter-UNIT-TESTS test delimiter-LINTERS lint delimiter-FINISH ## Runs multiple targets, audit, lint and test

.PHONY: audit
audit: audit-go audit-node

.PHONY: audit-go
audit-go: ## Runs checks for security vulnerabilities on dependencies (including transient ones)
	go list -json -m all | nancy sleuth

.PHONY: audit-node
audit-node: 
	$(NPM) audit

.PHONY: build
build: build-go build-node ## Builds the Go binary and the frontend assets

.PHONY: build-go
build-go: 
	go build -tags 'production' $(LDFLAGS) -o $(BINPATH)/dp-renderer

.PHONY: build-node
build-node: 
	$(NPM) install --unsafe-perm
	$(NPM) run build


.PHONY: convey
convey: ## Runs unit test suite and outputs results on http://127.0.0.1:8080/
	goconvey ./...

.PHONY: debug
debug: ## Runs a web server to serve the css and js files
	$(NPM) install && $(NPM) run dev

.PHONY: delimiter-%
delimiter-%:
	@echo '===================${GREEN} $* ${RESET}==================='

.PHONY: fmt
fmt: ## Run Go formatting on code
	go fmt ./...

.PHONY: lint
lint: lint-go lint-js ## Run all linters

.PHONY: lint-go
lint-go: ## Run Go linters
	golangci-lint run ./...

.PHONY: lint-js
lint-js: ## Run JS linters
	$(NPM) run lint

.PHONY: test
test: ## Runs unit tests including checks for race conditions and returns coverage
	go test -race -cover -tags 'production' ./...

.PHONY: help
help: ## Show help page for list of make targets
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-20s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
