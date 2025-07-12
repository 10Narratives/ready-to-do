#!/bin/bash

set -euo pipefail

# Пути
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(dirname "$SCRIPT_DIR")"
VENV_DIR="$SCRIPT_DIR/../.py-venv-protoc"
GO_BIN="$HOME/go/bin"
PLUGIN_PATH_GO="$GO_BIN/protoc-gen-go"
PLUGIN_PATH_SWAGGER="$GO_BIN/protoc-gen-openapiv2"
GOOGLE_APIS_DIR="$PROJECT_DIR/vendor/google"

if [ ! -d "$VENV_DIR" ]; then
  echo "❌ Virtual environment not found. Run ./scripts/prepare.sh first."
  exit 1
fi
source "$VENV_DIR/bin/activate"

if [ ! -f "$PLUGIN_PATH_GO" ]; then
  echo "❌ protoc-gen-go not found at $PLUGIN_PATH_GO"
  echo "💡 Run ./scripts/prepare.sh to install it."
  exit 1
fi

if [ ! -f "$PLUGIN_PATH_SWAGGER" ]; then
  echo "❌ protoc-gen-openapiv2 not found at $PLUGIN_PATH_SWAGGER"
  echo "💡 Run ./scripts/prepare.sh to install it."
  exit 1
fi

LANGUAGES="go python swagger"
PACKAGES="$(find "$PROJECT_DIR/proto" -mindepth 1 -maxdepth 1 -type d)"
GEN_ROOT_DIR="$PROJECT_DIR/gen"

while [[ "$#" -gt 0 ]]; do
  case $1 in
    --lang) LANGUAGES="$2"; shift ;;
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
      --go_out="$GEN_DIR" \
      --go_opt=paths=source_relative \
      --plugin=protoc-gen-go="$PLUGIN_PATH_GO" \
      "$file"
  done

  echo "[Go] ✅ Generated into $GEN_DIR"
}

generate_python() {
  GEN_DIR="$GEN_ROOT_DIR/python"
  rm -rf "$GEN_DIR"
  mkdir -p "$GEN_DIR"

  for file in $(get_proto_files); do
    protoc \
      --proto_path="$PROJECT_DIR" \
      --proto_path="$GOOGLE_APIS_DIR" \
      --python_out="$GEN_DIR" \
      "$file"
  done

  echo "[Python] ✅ Generated into $GEN_DIR"
}

generate_swagger() {
  GEN_DIR="$GEN_ROOT_DIR/swagger"
  rm -rf "$GEN_DIR"
  mkdir -p "$GEN_DIR"

  for file in $(get_proto_files); do
    protoc \
      --proto_path="$PROJECT_DIR" \
      --proto_path="$GOOGLE_APIS_DIR" \
      --plugin=protoc-gen-openapiv2="$PLUGIN_PATH_SWAGGER" \
      --openapiv2_out="$GEN_DIR" \
      --openapiv2_opt=logtostderr=true \
      "$file"
  done

  echo "[Swagger] ✅ Generated into $GEN_DIR"
}

rm -rf "$GEN_ROOT_DIR"
mkdir -p "$GEN_ROOT_DIR"

for lang in $LANGUAGES; do
  case "$lang" in
    go) generate_go ;;
    python) generate_python ;;
    swagger) generate_swagger ;;
    *)
      echo "❌ Unsupported language: $lang"
      exit 1
      ;;
  esac
done

echo "🎉 All files generated successfully!"