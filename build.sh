#!/bin/bash

# Fail on any error or undefined variable
set -e -o pipefail -u

SCRIPT_FILE="$(basename "$0")"
# NOTE: readlink will not work in OSX
SCRIPT_DIR="$(dirname "$(readlink -f "$0")")"

sudo docker build --target builder -t advent-of-code-2023-builder:local-build .

sudo docker build -t advent-of-code-2023:local-build .
