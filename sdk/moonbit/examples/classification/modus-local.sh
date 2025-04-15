#!/bin/bash -ex

# If no arguments are provided, set default to "dev"
args=("${@:-"dev"}")

export MODUS_DB=postgresql://postgres:postgres@localhost:5433/my-runtime-db?sslmode=disable

export MODUS_DEBUG=true
export MODUS_TRACE=true
../../../../cli/bin/modus.js "${args[@]}"
