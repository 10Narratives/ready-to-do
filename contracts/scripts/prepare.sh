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
PLUGIN_PATH="$GO_BIN/protoc-gen-go"

if ! command -v protoc-gen-go &> /dev/null; then
  echo "protoc-gen-go not found. Installing..."
  if ! command -v go &> /dev/null; then
    echo "❌ Go not found. Please install Go first: https://go.dev/dl/ "
    exit 1
  fi

  mkdir -p "$GO_BIN"
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  echo "✅ protoc-gen-go installed to $PLUGIN_PATH"
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