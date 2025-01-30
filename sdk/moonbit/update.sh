#!/bin/bash -ex
moon add moonbitlang/x
moon add gmlewis/base64
moon update && moon install && rm -rf target
moon fmt
moon test --target wasm-gc
moon test --target js
moon test --target native
