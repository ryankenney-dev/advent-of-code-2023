#!/bin/bash

# Fail on any error or undefined variable
set -e -o pipefail -u

sudo docker run --rm -p 127.0.0.1:8080:8080 -it --name advent-of-code-2023 advent-of-code-2023:local-build
