#!/bin/bash -ex

# Load environment variables
source .env.development.local

# If no arguments are provided, set default to "dev"
args=("${@:-"dev"}")

# export MODUS_DEBUG=true
export MODUS_TRACE=true
../../../../cli/bin/modus.js "${args[@]}"
