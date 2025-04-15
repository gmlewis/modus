#!/bin/bash

rm -rf .modusdb build target

PROJECTDIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
pushd ../../../../sdk/go/tools/modus-moonbit-build >/dev/null || exit
MODUS_DEBUG=true
MODUS_TRACE=true
go run . "${PROJECTDIR}"
exit_code=$?
popd >/dev/null || exit

if command -v wasm2wat >/dev/null 2>&1; then
  # If wasm2wat is available, run the command
  wasm2wat build/testdata.wasm > build/testdata.wat
fi

exit "${exit_code}"
