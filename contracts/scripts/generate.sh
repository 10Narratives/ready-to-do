#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
VENV_DIR="$SCRIPT_DIR/../.py-venv-protoc"
GO_BIN="$HOME/go/bin"

PLUGIN_PATHS=(
  "go:$GO_BIN/protoc-gen-go"
  "go-grpc:$GO_BIN/protoc-gen-go-grpc"
  "gateway:$GO_BIN/protoc-gen-grpc-gateway"
  "swagger:$GO_BIN/protoc-gen-openapiv2"
  "validate:$GO_BIN/protoc-gen-validate"
)

if [ ! -d "$VENV_DIR" ]; then
  echo "❌ Virtual environment not found. Run ./scripts/prepare.sh first."
  exit 1
fi
source "$VENV_DIR/bin/activate"

for plugin in "${PLUGIN_PATHS[@]}"; do
  name=${plugin%%:*}
  path=${plugin#*:}
  if [ ! -f "$path" ]; then
    echo "❌ Plugin not found: protoc-gen-$name ($path)"
    echo "💡 Run ./scripts/prepare.sh to install it."
    exit 1
  fi
done

GOOGLE_APIS_DIR="$PROJECT_DIR/vendor/google"
VALIDATE_DIR="$PROJECT_DIR/vendor/protoc-gen-validate"

for dir in "$GOOGLE_APIS_DIR" "$VALIDATE_DIR"; do
  if [ ! -d "$dir" ]; then
    echo "❌ Dependency not found: $dir"
    echo "💡 Run ./scripts/prepare.sh to install it."
    exit 1
  fi
done

LANGUAGES="go python"
FEATURES="swagger grpc-gateway"
PACKAGES="$(find "$PROJECT_DIR/proto" -mindepth 1 -maxdepth 1 -type d)"
GEN_ROOT_DIR="$PROJECT_DIR/gen"

while [[ "$#" -gt 0 ]]; do
  case $1 in
    --lang) LANGUAGES="$2"; shift ;;
    --feature) FEATURES="$2"; shift ;;
    --pkg) PACKAGES="$2"; shift ;;
    *) echo "Unknown parameter passed: $1"; exit 1 ;;
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

  for file in $(get_proto_files); do
    protoc \
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
      "$file"
  done

  echo "[Go] ✅ Generated into $GEN_DIR (with validation)"
}

generate_python() {
  GEN_DIR="$GEN_ROOT_DIR/python"
  rm -rf "$GEN_DIR"
  mkdir -p "$GEN_DIR"

  for file in $(get_proto_files); do
    python -m grpc_tools.protoc \
      --proto_path="$PROJECT_DIR" \
      --proto_path="$GOOGLE_APIS_DIR" \
      --proto_path="$VALIDATE_DIR" \
      --python_out="$GEN_DIR" \
      --grpc_python_out="$GEN_DIR" \
      "$file"
  done

  echo "[Python] ✅ Generated messages and gRPC services into $GEN_DIR"
}

generate_swagger() {
  GEN_DIR="$GEN_ROOT_DIR/swagger"
  rm -rf "$GEN_DIR"
  mkdir -p "$GEN_DIR"

  for file in $(get_proto_files); do
    protoc \
      --proto_path="$PROJECT_DIR" \
      --proto_path="$GOOGLE_APIS_DIR" \
      --proto_path="$VALIDATE_DIR" \
      --plugin=protoc-gen-openapiv2="$GO_BIN/protoc-gen-openapiv2" \
      --openapiv2_out="$GEN_DIR" \
      --openapiv2_opt=logtostderr=true \
      "$file"
  done

  echo "[Swagger] ✅ Generated into $GEN_DIR"
}

generate_grpc_gateway() {
  GEN_DIR="$GEN_ROOT_DIR/grpc-gateway"
  rm -rf "$GEN_DIR"
  mkdir -p "$GEN_DIR"

  for file in $(get_proto_files); do
    protoc \
      --proto_path="$PROJECT_DIR" \
      --proto_path="$GOOGLE_APIS_DIR" \
      --proto_path="$VALIDATE_DIR" \
      --grpc-gateway_out="$GEN_DIR" \
      --grpc-gateway_opt=logtostderr=true \
      --grpc-gateway_opt=paths=source_relative \
      --grpc-gateway_opt=generate_unbound_methods=true \
      --plugin=protoc-gen-grpc-gateway="$GO_BIN/protoc-gen-grpc-gateway" \
      "$file"
  done

  echo "[gRPC-Gateway] ✅ Generated into $GEN_DIR"
}

for lang in $LANGUAGES; do
  case "$lang" in
    go) generate_go ;;
    python) generate_python ;;
    *)
      echo "❌ Unsupported language: $lang"
      exit 1
      ;;
  esac
done

for feature in $FEATURES; do
  case "$feature" in
    swagger) generate_swagger ;;
    grpc-gateway) generate_grpc_gateway ;;
    *)
      echo "❌ Unsupported feature: $feature"
      exit 1
      ;;
  esac
done

echo "🎉 All files generated successfully!"