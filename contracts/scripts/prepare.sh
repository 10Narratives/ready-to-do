#!/bin/bash

set -euo pipefail

echo "🔧 Preparing environment..."

if ! command -v protoc &> /dev/null; then
  echo "protobuf (protoc) not found. Installing..."
  case "$(uname -s)" in
    Linux)
      if grep -q "Arch" /etc/os-release; then
        sudo pacman -Sy --noconfirm protobuf
      else
        sudo apt-get update && sudo apt-get install -y protobuf-compiler
      fi
      ;;
    Darwin)
      brew install protobuf
      ;;
    *)
      echo "Unsupported OS for automatic protoc installation."
      exit 1
      ;;
  esac
else
  echo "protoc is already installed"
fi

GO_BIN="$HOME/go/bin"

PLUGIN_PATH_GO="$GO_BIN/protoc-gen-go"
PLUGIN_PATH_GO_GRPC="$GO_BIN/protoc-gen-go-grpc"
PLUGIN_PATH_GATEWAY="$GO_BIN/protoc-gen-grpc-gateway"
PLUGIN_PATH_SWAGGER="$GO_BIN/protoc-gen-openapiv2"

if ! command -v protoc-gen-go &> /dev/null; then
  echo "protoc-gen-go not found. Installing..."
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
fi

if ! command -v protoc-gen-go-grpc &> /dev/null; then
  echo "protoc-gen-go-grpc not found. Installing..."
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
fi

if ! command -v protoc-gen-grpc-gateway &> /dev/null; then
  echo "protoc-gen-grpc-gateway not found. Installing..."
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
fi

if ! command -v protoc-gen-openapiv2 &> /dev/null; then
  echo "protoc-gen-openapiv2 not found. Installing..."
  go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
fi

VENDOR_DIR="vendor"
GOOGLE_APIS_DIR="$VENDOR_DIR/google"

if [ ! -d "$GOOGLE_APIS_DIR" ]; then
  echo "📦 Cloning googleapis into $GOOGLE_APIS_DIR..."
  git clone --depth=1 https://github.com/googleapis/googleapis.git  "$GOOGLE_APIS_DIR"
else
  echo "📦 googleapis already exists in $GOOGLE_APIS_DIR"
fi

VENV_DIR=".py-venv-protoc"
rm -rf "$VENV_DIR"
python -m venv "$VENV_DIR"
source "$VENV_DIR/bin/activate"

echo "Installing Python protobuf tools..."
pip install --no-cache-dir protobuf grpcio-tools

export PATH="$GO_BIN:$PATH"

echo "✅ Environment is ready!"