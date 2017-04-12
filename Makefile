.PHONY: help tidy test test-crypto test-all lint build-gateway build-all coverage bootstrap loc docker-up docker-down

GO ?= go
SERVICES := $(shell ls apps)

help:
	@echo "SentinelVault targets:"
	@echo "  make tidy            - go mod tidy"
	@echo "  make test            - run core package tests"
	@echo "  make test-all        - run all tests"
	@echo "  make lint            - golangci-lint"
	@echo "  make build-gateway   - build api-gateway"
	@echo "  make bootstrap       - verify toolchain"
	@echo "  make loc             - count lines of code"
	@echo "  make docker-up       - start local stack"

tidy:
	$(GO) mod tidy

test:
	$(GO) test ./internal/common/... ./internal/crypto/... ./internal/config/...

test-all:
	$(GO) test ./...

lint:
	golangci-lint run ./...

build-gateway:
	$(GO) build -o bin/api-gateway ./apps/api-gateway/cmd/

build-all:
	@for svc in $(SERVICES); do \
		echo "building $$svc"; \
		$(GO) build -o bin/$$svc ./apps/$$svc/cmd/ || exit 1; \
	done

coverage:
	$(GO) test -coverprofile=coverage.out ./internal/common/... ./internal/crypto/... ./internal/config/...
	$(GO) tool cover -func=coverage.out

bootstrap:
	./scripts/bootstrap-check.sh

loc:
	./scripts/loc.sh

docker-up:
	docker compose -f deploy/docker/docker-compose.yml up -d

docker-down:
	docker compose -f deploy/docker/docker-compose.yml down

generate:
	python3 scripts/generate_services.py
