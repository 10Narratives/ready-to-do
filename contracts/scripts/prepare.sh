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


PLUGINS=(
  "protoc-gen-go:google.golang.org/protobuf/cmd/protoc-gen-go@latest"
  "protoc-gen-go-grpc:google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest"
  "protoc-gen-grpc-gateway:github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest"
  "protoc-gen-openapiv2:github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest"
  "protoc-gen-validate:github.com/envoyproxy/protoc-gen-validate@latest"
)

for plugin in "${PLUGINS[@]}"; do
  name=${plugin%%:*}
  pkg=${plugin#*:}
  if ! command -v $name &> /dev/null; then
    echo "Installing $name..."
    go install $pkg
  else
    echo "$name is already installed"
  fi
done

VENDOR_DIR="vendor"
mkdir -p "$VENDOR_DIR"

GOOGLE_APIS_DIR="$VENDOR_DIR/google"
if [ ! -d "$GOOGLE_APIS_DIR" ]; then
  echo "📦 Cloning googleapis into $GOOGLE_APIS_DIR..."
  git clone --depth=1 https://github.com/googleapis/googleapis.git "$GOOGLE_APIS_DIR"
else
  echo "📦 googleapis already exists in $GOOGLE_APIS_DIR"
fi

VALIDATE_DIR="$VENDOR_DIR/protoc-gen-validate"
if [ ! -d "$VALIDATE_DIR" ]; then
  echo "📦 Cloning protoc-gen-validate into $VALIDATE_DIR..."
  git clone --depth=1 https://github.com/envoyproxy/protoc-gen-validate.git "$VALIDATE_DIR"
else
  echo "📦 protoc-gen-validate already exists in $VALIDATE_DIR"
fi

VENV_DIR=".py-venv-protoc"
rm -rf "$VENV_DIR"
python -m venv "$VENV_DIR"
source "$VENV_DIR/bin/activate"

echo "Installing Python protobuf tools..."
pip install --no-cache-dir protobuf grpcio-tools

export PATH="$GO_BIN:$PATH"

echo "✅ Environment is ready!"