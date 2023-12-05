#!/bin/bash

# Fail on any error or undefined variable
set -e -o pipefail -u

SCRIPT_FILE="$(basename "$0")"
# NOTE: readlink will not work in OSX
SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

bash "${SCRIPT_DIR}/build.sh"

bash "${SCRIPT_DIR}/test.sh"
