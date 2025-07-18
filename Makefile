BIN := dist/hrs-go

.PHONY: help
help:
	@grep -E '^[0-9a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":[^:]*?## "}; {printf "%-20s %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the distributable
	go build -o ${BIN} cmd/hrs/main.go

.PHONY: test
test: ## Run short tests
	go tool gotestsum --format=testname -- -short -count=1 ./...

.PHONY: e2e
e2e: build ## Run e2e tests
	go tool gotestsum --format=testname -- --count=1 -run ^TestE2E$$ github.com/istonikula/hrs-go/cmd/hrs

.PHONY: clean
clean: ## Clean intermediate build products and remove distributable
	go clean
	$(RM) -f ${BIN}

.PHONY: lint
lint: ## Run lint
	go tool golangci-lint run --timeout 5m
