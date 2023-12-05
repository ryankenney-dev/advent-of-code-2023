#!/bin/bash

# Fail on any error or undefined variable
set -e -o pipefail -u

sudo docker run --rm -it --name advent-of-code-2023-test advent-of-code-2023-builder:local-build go test ./...
