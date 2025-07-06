#!/bin/bash

PROTO_DIR="proto/"
PROTO_OUT_DIR="pkg/"
GOOGLEAPIS="/tmp/googleapis"
PGV_DIR="/tmp/protoc-gen-validate"

clean() {
    echo "Cleaning generated protobuf files..."
    rm -rf "$PROTO_OUT_DIR"
    echo "Cleanup complete."
}

install_deps() {
    echo "Installing dependencies..."
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/envoyproxy/protoc-gen-validate@latest
    echo "Dependencies installed successfully."
}

clone_repos() {
    echo "Ensuring required repositories are cloned..."
    if [ ! -d "$GOOGLEAPIS" ]; then
        echo "Cloning googleapis..."
        git clone https://github.com/googleapis/googleapis.git  "$GOOGLEAPIS"
    fi
    if [ ! -d "$PGV_DIR" ]; then
        echo "Cloning protoc-gen-validate..."
        git clone https://github.com/envoyproxy/protoc-gen-validate.git  "$PGV_DIR"
    fi
    echo "Repositories are ready."
}

generate_code() {
    clean
    install_deps
    clone_repos

    echo "Generating protobuf, gRPC, and validation code..."

    mkdir -p "$PROTO_OUT_DIR"

    if [ "$#" -eq 0 ]; then
        proto_files=$(find "$PROTO_DIR" -name "*.proto")
    else
        proto_files=()
        for package in "$@"; do
            proto_files+=($(find "$PROTO_DIR/$package" -name "*.proto" 2>/dev/null))
        done
    fi

    if [ ${#proto_files[@]} -eq 0 ]; then
        echo "No .proto files found for the specified packages. Exiting..."
        exit 1
    fi

    for proto_file in "${proto_files[@]}"; do
        echo "Processing $proto_file..."
        protoc -I"$PROTO_DIR" \
               -I"$GOOGLEAPIS" \
               -I"$PGV_DIR" \
               "$proto_file" \
               --go_out="$PROTO_OUT_DIR" --go_opt=paths=source_relative \
               --go-grpc_out="$PROTO_OUT_DIR" --go-grpc_opt=paths=source_relative \
               --validate_out="lang=go,paths=source_relative:$PROTO_OUT_DIR"
    done

    echo "Code generation complete."
}

run_app() {
    echo "Starting the worker application..."
    go run cmd/worker/main.go --config config/worker.yaml
    echo "Worker application started."
}

show_help() {
    echo "Usage: $0 [command] [arguments]"
    echo ""
    echo "Available commands:"
    echo "  clean      - Remove all generated protobuf files and temporary directories."
    echo "  install    - Install required Go tools (protoc-gen-go, protoc-gen-go-grpc, protoc-gen-validate)."
    echo "  deps       - Clone required repositories (googleapis and protoc-gen-validate) if they don't exist."
    echo "  generate   - Generate Go protobuf, gRPC, and validation code from .proto files."
    echo "               Usage: generate [package1] [package2] ..."
    echo "               If no packages are specified, all .proto files will be processed."
    echo "  run        - Run the worker application using the specified configuration file."
    echo "  help       - Display this help message."
}

case "$1" in
    clean)
        clean
        ;;
    install)
        install_deps
        ;;
    deps)
        clone_repos
        ;;
    generate)
        shift
        generate_code "$@"
        ;;
    run)
        run_app
        ;;
    help)
        show_help
        ;;
    *)
        echo "Unknown command: $1"
        show_help
        exit 1
        ;;
esac