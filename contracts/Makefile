override LANG = go python
override FEATURE = swagger grpc-gateway
override PKG = $(shell find proto -mindepth 1 -maxdepth 1 -type d)

.PHONY: all prepare generate help

all: help

help:
	@echo ""
	@echo "📌 Available Make Targets:"
	@echo "-----------------------------------------------"
	@echo ""
	@echo "  🛠️  make prepare"
	@echo "      Install all required dependencies."
	@echo ""
	@echo "  🧱 make generate [LANG=\"go python\"] [FEATURE=\"swagger grpc-gateway\"] [PKG=\"proto/package1 proto/package2\"]"
	@echo "      Generate code and API artifacts."
	@echo ""
	@echo "      LANG    — target languages, default: go python"
	@echo "      FEATURE — additional API/tooling support, default: swagger grpc-gateway"
	@echo "      PKG     — proto directories to process, default: all subdirectories in proto/"
	@echo ""
	@echo "💡 Examples:"
	@echo "   make generate LANG=go FEATURE=grpc-gateway PKG=proto/tracker"
	@echo "   make generate LANG=\"go python\" FEATURE=\"swagger grpc-gateway\""
	@echo ""

prepare:
	./scripts/prepare.sh

generate:
	./scripts/generate.sh --lang "$(LANG)" --feature "$(FEATURE)" --pkg "$(PKG)"