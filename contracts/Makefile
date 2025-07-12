LANGUAGE ?= go python swagger grpc-gateway
PKG ?= $(shell find proto -mindepth 1 -maxdepth 1 -type d)

.PHONY: all prepare generate help

all: help

help:
	@echo ""
	@echo "📌 Available Make Targets:"
	@echo "-----------------------------------------------"
	@echo ""
	@echo "  🛠️  make prepare"
	@echo "      Install all required dependencies for code generation."
	@echo ""
	@echo "  🧱 make generate [LANGUAGE=\"go python swagger grpc-gateway\"] [PKG=\"proto/package1 proto/package2\"]"
	@echo "      Generate protocol buffer artifacts and/or gRPC gateway + Swagger UI"
	@echo ""
	@echo "      LANGUAGE — target languages, default: go python swagger grpc-gateway"
	@echo "      PKG      — proto directories to process, default: all subdirectories in proto/"
	@echo ""
	@echo "💡 Examples:"
	@echo "   make generate LANGUAGE=grpc-gateway PKG=proto/tracker"
	@echo "   make generate LANGUAGE=\"go grpc-gateway\" PKG=\"proto/example1 proto/example2\""
	@echo ""

prepare:
	./scripts/prepare.sh

generate:
	./scripts/generate.sh --lang "$(LANGUAGE)" --pkg "$(PKG)"