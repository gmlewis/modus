#!/bin/bash -ex

# If no arguments are provided, set default to "dev"
args=("${@:-"dev"}")

../../../../cli/bin/modus.js "${args[@]}"
