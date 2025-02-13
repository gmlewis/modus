#!/bin/bash -ex
rm -rf ffi interface pkg world
wit-bindgen moonbit --derive-show --derive-eq --derive-error --gen-dir pkg .
mv interface/modus/core/* interface
rm -rf interface/modus
rm -rf pkg/world* world moon.mod.json
rep 'modus/core/' 'gmlewis/modus/wit/' pkg/moon.pkg.json interface/*/moon.pkg.json
rep 'modus:core/' 'modus_' interface/*/ffi.mbt
rep -- '-client' '_client' interface/*/ffi.mbt
rep 'derive(Show, Eq)' 'derive(Show, Eq, FromJson, ToJson)' interface/*/top.mbt
# Rename the *.mbt files to *_wasm.mbt:
find . -name "*.mbt" | while read -r f; do mv "$f" "${f%.mbt}_wasm.mbt"; done
# Now recreate the moon.pkg.json files:
find . -name "*_wasm.mbt" -exec dirname {} \; | sort -u | while read -r dir; do
  echo "{\"targets\":{$(find "$dir" -maxdepth 1 -name "*_wasm.mbt" -exec basename {} \; | sed 's/.*/"&":["wasm"]/' | paste -sd, -)}}" > "$dir/moon.pkg.json"
done
moon fmt
