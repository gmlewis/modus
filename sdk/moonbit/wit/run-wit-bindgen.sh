#!/bin/bash -ex
rm -rf ffi interface pkg world
wit-bindgen moonbit --derive-show --derive-eq --derive-error --gen-dir pkg .
mv interface/modus/core/* interface
rm -rf pkg/world* world moon.mod.json
rep 'modus/core/' 'gmlewis/modus/wit/' pkg/moon.pkg.json interface/*/moon.pkg.json
rep 'modus:core/' 'modus_' interface/*/ffi.mbt
rep -- '-client' '_client' interface/*/ffi.mbt
moon fmt
