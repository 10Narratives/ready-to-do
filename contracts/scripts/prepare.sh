#!/bin/bash

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
source "$SCRIPT_DIR/logging.sh"
init_logging "[ProtoEnv]"

VENDOR_DIR="vendor"
VENV_DIR=".py-venv-protoc"
GO_BIN="$HOME/go/bin"

run_pip_install() {
  local cmd="$1"
  local description="$2"
  local log_file

  log_file=$(mktemp)
  log_step "$description"

  {
    eval "$cmd" 2>&1 | while IFS= read -r line; do
      if [[ "$line" =~ Collecting|Downloading|Already|Successfully ]]; then
        echo -e "${DIM}${line}${NC}"
      fi
    done
    exit "${PIPESTATUS[0]}"
  } > >(tee "$log_file")

  if [ ${PIPESTATUS[0]} -ne 0 ]; then
    log_error "Failed to $description"
    echo -e "${DIM}Full log: $log_file${NC}" >&2
    exit 1
  fi

  rm -f "$log_file"
}

log_header "Preparing Protocol Buffers Environment"

if ! command -v protoc &>/dev/null; then
  case "$(uname -s)" in
  Linux)
    if grep -q "Arch" /etc/os-release; then
      run_command "sudo pacman -Sy --noconfirm protobuf" "Install protoc (Arch Linux)"
    else
      run_command "sudo apt-get update && sudo apt-get install -y protobuf-compiler" "Install protoc (Debian/Ubuntu)"
    fi
    ;;
  Darwin)
    run_command "brew install protobuf" "Install protoc (macOS)"
    ;;
  *)
    log_error "Unsupported OS for automatic protoc installation"
    exit 1
    ;;
  esac
else
  log_success "protoc is already installed"
fi

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
  if ! command -v $name &>/dev/null; then
    run_command "go install $pkg" "Install $name"
  else
    log_success "$name is already installed"
  fi
done

run_command "mkdir -p \"$VENDOR_DIR\"" "Create vendor directory"

GOOGLE_APIS_DIR="$VENDOR_DIR/google"
if [ ! -d "$GOOGLE_APIS_DIR" ]; then
  run_command "git clone --depth=1 https://github.com/googleapis/googleapis.git \"$GOOGLE_APIS_DIR\"" "Clone googleapis"
else
  log_success "googleapis already exists"
fi

VALIDATE_DIR="$VENDOR_DIR/protoc-gen-validate"
if [ ! -d "$VALIDATE_DIR" ]; then
  run_command "git clone --depth=1 https://github.com/envoyproxy/protoc-gen-validate.git \"$VALIDATE_DIR\"" "Clone protoc-gen-validate"
else
  log_success "protoc-gen-validate already exists"
fi

run_command "rm -rf \"$VENV_DIR\" && python -m venv \"$VENV_DIR\"" "Create Python virtual environment"

source "$VENV_DIR/bin/activate"

run_pip_install "pip install --no-cache-dir protobuf grpcio-tools" "Install Python protobuf tools"

export PATH="$GO_BIN:$PATH"
log_step "Added \$GO_BIN to PATH: $GO_BIN"

echo -e "\n${BOLD}${GREEN}${ICON_ROCKET} Environment preparation completed successfully! ${NC}"
echo -e "${DIM}You can now run the generation script with:${NC}"
echo -e "${BOLD}./scripts/generate.sh${NC}\n"
