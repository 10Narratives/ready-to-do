#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
VENV_DIR="$SCRIPT_DIR/../.py-venv-protoc"
GO_BIN="$HOME/go/bin"

source "$SCRIPT_DIR/logging.sh"
init_logging "[ProtoGen]"

PLUGIN_PATHS=(
  "go:$GO_BIN/protoc-gen-go"
  "go-grpc:$GO_BIN/protoc-gen-go-grpc"
  "gateway:$GO_BIN/protoc-gen-grpc-gateway"
  "swagger:$GO_BIN/protoc-gen-openapiv2"
  "validate:$GO_BIN/protoc-gen-validate"
)

# Check virtual environment
if [ ! -d "$VENV_DIR" ]; then
    log_error "Virtual environment not found at $VENV_DIR"
    echo -e "${DIM}Run ./scripts/prepare.sh first${NC}" >&2
    exit 1
fi
source "$VENV_DIR/bin/activate"

# Check plugins
for plugin in "${PLUGIN_PATHS[@]}"; do
    name=${plugin%%:*}
    path=${plugin#*:}
    if [ ! -f "$path" ]; then
        log_error "Plugin not found: protoc-gen-$name"
        echo -e "${DIM}Expected at: $path\nRun ./scripts/prepare.sh to install it${NC}" >&2
        exit 1
    fi
done

# Check dependencies
GOOGLE_APIS_DIR="$PROJECT_DIR/vendor/google"
VALIDATE_DIR="$PROJECT_DIR/vendor/protoc-gen-validate"

for dir in "$GOOGLE_APIS_DIR" "$VALIDATE_DIR"; do
    if [ ! -d "$dir" ]; then
        log_error "Dependency not found: $dir"
        echo -e "${DIM}Run ./scripts/prepare.sh to install it${NC}" >&2
        exit 1
    fi
done

# Default parameters
LANGUAGES="go python"
FEATURES="swagger grpc-gateway"
PACKAGES="$(find "$PROJECT_DIR/proto" -mindepth 1 -maxdepth 1 -type d)"
GEN_ROOT_DIR="$PROJECT_DIR/gen"

# Parse arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --lang) LANGUAGES="$2"; shift ;;
        --feature) FEATURES="$2"; shift ;;
        --pkg) PACKAGES="$2"; shift ;;
        *) 
            log_error "Unknown parameter passed: $1"
            exit 1
            ;;
    esac
    shift
done

get_proto_files() {
    for dir in $PACKAGES; do
        find "$dir" -type f -name "*.proto"
    done
}

generate_go() {
    GEN_DIR="$GEN_ROOT_DIR/go"
    rm -rf "$GEN_DIR"
    mkdir -p "$GEN_DIR"
    local success=true
    local errors=""

    for file in $(get_proto_files); do
        if ! protoc \
            --proto_path="$PROJECT_DIR" \
            --proto_path="$GOOGLE_APIS_DIR" \
            --proto_path="$VALIDATE_DIR" \
            --go_out="$GEN_DIR" \
            --go_opt=paths=source_relative \
            --go-grpc_out="$GEN_DIR" \
            --go-grpc_opt=paths=source_relative \
            --validate_out="lang=go:$GEN_DIR" \
            --validate_opt=paths=source_relative \
            --plugin=protoc-gen-go="$GO_BIN/protoc-gen-go" \
            --plugin=protoc-gen-go-grpc="$GO_BIN/protoc-gen-go-grpc" \
            --plugin=protoc-gen-validate="$GO_BIN/protoc-gen-validate" \
            "$file" 2> >(grep -v "but not used" >&2); then
            success=false
            errors+="$(basename "$file") "
        fi
    done

    if $success; then
        generation_results+=("Go|$GEN_DIR|${GREEN}✓ Success${NC}")
    else
        generation_results+=("Go|$GEN_DIR|${RED}✗ Failed: ${errors}${NC}")
    fi
}

generate_python() {
    GEN_DIR="$GEN_ROOT_DIR/python"
    rm -rf "$GEN_DIR"
    mkdir -p "$GEN_DIR"
    local success=true
    local errors=""

    for file in $(get_proto_files); do
        if ! python -m grpc_tools.protoc \
            --proto_path="$PROJECT_DIR" \
            --proto_path="$GOOGLE_APIS_DIR" \
            --proto_path="$VALIDATE_DIR" \
            --python_out="$GEN_DIR" \
            --grpc_python_out="$GEN_DIR" \
            "$file" 2>&1; then
            success=false
            errors+="$(basename "$file") "
        fi
    done

    if $success; then
        generation_results+=("Python|$GEN_DIR|${GREEN}✓ Success${NC}")
    else
        generation_results+=("Python|$GEN_DIR|${RED}✗ Failed: ${errors}${NC}")
    fi
}

generate_swagger() {
    GEN_DIR="$GEN_ROOT_DIR/swagger"
    rm -rf "$GEN_DIR"
    mkdir -p "$GEN_DIR"
    local success=true
    local errors=""

    for file in $(get_proto_files); do
        if ! protoc \
            --proto_path="$PROJECT_DIR" \
            --proto_path="$GOOGLE_APIS_DIR" \
            --proto_path="$VALIDATE_DIR" \
            --plugin=protoc-gen-openapiv2="$GO_BIN/protoc-gen-openapiv2" \
            --openapiv2_out="$GEN_DIR" \
            --openapiv2_opt=logtostderr=true \
            "$file" 2>&1; then
            success=false
            errors+="$(basename "$file") "
        fi
    done

    if $success; then
        generation_results+=("Swagger|$GEN_DIR|${GREEN}✓ Success${NC}")
    else
        generation_results+=("Swagger|$GEN_DIR|${RED}✗ Failed: ${errors}${NC}")
    fi
}

generate_grpc_gateway() {
    GEN_DIR="$GEN_ROOT_DIR/go"
    local success=true
    local errors=""

    for file in $(get_proto_files); do
        if ! protoc \
            --proto_path="$PROJECT_DIR" \
            --proto_path="$GOOGLE_APIS_DIR" \
            --proto_path="$VALIDATE_DIR" \
            --grpc-gateway_out="$GEN_DIR" \
            --grpc-gateway_opt=logtostderr=true \
            --grpc-gateway_opt=paths=source_relative \
            --grpc-gateway_opt=generate_unbound_methods=true \
            --plugin=protoc-gen-grpc-gateway="$GO_BIN/protoc-gen-grpc-gateway" \
            "$file" 2>&1; then
            success=false
            errors+="$(basename "$file") "
        fi
    done

    if $success; then
        generation_results+=("gRPC Gateway|$GEN_DIR|${GREEN}✓ Success${NC}")
    else
        generation_results+=("gRPC Gateway|$GEN_DIR|${RED}✗ Failed: ${errors}${NC}")
    fi
}

main() {
    local generation_results=()
    
    log_header "Starting code generation"
    echo -e "${DIM}Languages: ${LANGUAGES}\nFeatures: ${FEATURES}\nPackages: ${PACKAGES}${NC}"

    for lang in $LANGUAGES; do
        case "$lang" in
            go) generate_go ;;
            python) generate_python ;;
            *)
                log_error "Unsupported language: $lang"
                exit 1
                ;;
        esac
    done

    for feature in $FEATURES; do
        case "$feature" in
            swagger) generate_swagger ;;
            grpc-gateway) generate_grpc_gateway ;;
            *)
                log_error "Unsupported feature: $feature"
                exit 1
                ;;
        esac
    done

    print_generation_table "${generation_results[@]}"

    # Check for failures
    for result in "${generation_results[@]}"; do
        if [[ "$result" == *"✗ Failed"* ]]; then
            echo -e "\n${BOLD}${RED}Some generations failed!${NC}"
            exit 1
        fi
    done

    echo -e "\n${BOLD}${GREEN}${ICON_ROCKET} All files generated successfully! ${NC}"
}

main "$@"