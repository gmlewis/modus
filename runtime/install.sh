#!/bin/bash -ex
go build
mv ./runtime ${HOME}/.modus/runtime/v0.16.3/modus_runtime
