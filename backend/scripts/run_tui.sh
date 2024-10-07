#!/bin/bash

set -e

SCRIPT_PATH="$(dirname "$(realpath "$0")")"

source "$SCRIPT_PATH/env.sh"

go run ../src/cmd/main.go