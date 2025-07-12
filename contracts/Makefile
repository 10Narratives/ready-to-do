override LANG = go python
override FEATURE = swagger grpc-gateway
override PKG = $(shell find proto -mindepth 1 -maxdepth 1 -type d)

.PHONY: all prepare generate help

.DEFAULT_GOAL := help

all: help

help:
	@echo "Usage: make <target> [options]"
	@echo ""
	@echo "Targets:"
	@echo "  prepare   Install required dependencies"
	@echo "  generate  Generate code and API artifacts"
	@echo "  help      Show this help message"
	@echo ""
	@echo "Options for generate:"
	@echo "  LANG      Target languages (space-separated, default: go python)"
	@echo "  FEATURE   Additional features (space-separated, default: swagger grpc-gateway)"
	@echo "  PKG       Proto package directories to process (default: all proto subdirectories)"
	@echo ""
	@echo "Examples:"
	@echo "  make generate LANG=go FEATURE=grpc-gateway PKG=proto/tracker"
	@echo "  make generate LANG=\"go python\" FEATURE=\"swagger grpc-gateway\""

prepare:
	@echo "Installing dependencies..."
	@./scripts/prepare.sh

generate:
	@echo "Generating artifacts..."
	@./scripts/generate.sh --lang "$(LANG)" --feature "$(FEATURE)" --pkg "$(PKG)"