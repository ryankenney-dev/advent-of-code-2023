#!/bin/bash

set -e -o pipefail -u

sudo docker build -t advent-of-code-2023:local-build .

sudo docker run --rm -p 127.0.0.1:8080:8080 -it advent-of-code-2023:local-build

