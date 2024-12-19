#!/bin/bash -ex
rm -rf ffi interface pkg world
wit-bindgen moonbit --derive-show --derive-eq --derive-error --gen-dir pkg .
rm -rf pkg/world* world
