#!/bin/bash

# --- Style Definitions ---
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
MAGENTA='\033[0;35m'
NC='\033[0m' # No Color
BOLD='\033[1m'
DIM='\033[2m'

# --- Icons ---
ICON_CHECK="${GREEN}✓${NC}"
ICON_CROSS="${RED}✗${NC}"
ICON_GEAR="${BLUE}⚙${NC}"
ICON_ROCKET="${GREEN}🚀${NC}"

# --- Logging Functions ---
init_logging() {
    LOG_PREFIX="${1:-[ProtoGen]}"
    export LOG_PREFIX
}

log_header() {
    echo -e "\n${BOLD}${MAGENTA}${ICON_GEAR} ${LOG_PREFIX} $1 ${NC}"
}

log_step() {
    echo -e "${ICON_GEAR} ${BOLD}${LOG_PREFIX}${NC} $1"
}

log_success() {
    echo -e "${ICON_CHECK} ${LOG_PREFIX} ${GREEN}$1${NC}"
}

log_error() {
    echo -e "${ICON_CROSS} ${LOG_PREFIX} ${RED}$1${NC}" >&2
}

print_generation_table() {
    local headers=("Language/Feature" "Output Directory" "Status")
    local rows=("$@")
    
    local col1_len=${#headers[0]}
    local col2_len=${#headers[1]}
    local col3_len=${#headers[2]}
    
    for row in "${rows[@]}"; do
        IFS='|' read -ra cols <<< "$row"

        local len1=$(echo -e "${cols[0]}" | sed -r "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" | wc -m)
        ((len1 > col1_len)) && col1_len=$len1
        

        local len2=$(echo -e "${cols[1]}" | sed -r "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" | wc -m)
        ((len2 > col2_len)) && col2_len=$len2
        
        col3_len=12
    done
    
    ((col1_len+=2))
    ((col2_len+=2))
    
    echo -e "\n${BOLD}${MAGENTA}Generation Results:${NC}"
    printf "%-${col1_len}s | %-${col2_len}s | %-${col3_len}s\n" "${headers[0]}" "${headers[1]}" "${headers[2]}"
    printf "%${col1_len}s | %${col2_len}s | %${col3_len}s\n" | tr ' ' '-'
    
    for row in "${rows[@]}"; do
        IFS='|' read -ra cols <<< "$row"
        printf "%-${col1_len}s | %-${col2_len}s | %-${col3_len}b\n" "${cols[0]}" "${cols[1]}" "${cols[2]}"
    done
}

run_command() {
    local cmd="$1"
    local description="$2"
    local log_file
    
    log_file=$(mktemp)
    log_step "$description"
    
    if eval "$cmd" > "$log_file" 2>&1; then
        log_success "$description completed"
        if [ -s "$log_file" ]; then
            echo -e "${DIM}$(sed 's/^/  /' "$log_file")${NC}"
        fi
        return 0
    else
        log_error "Failed to $description"
        echo -e "${DIM}$(sed 's/^/  /' "$log_file")${NC}" >&2
        rm -f "$log_file"
        return 1
    fi
    
    rm -f "$log_file"
}