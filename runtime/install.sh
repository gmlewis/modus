#!/bin/bash -ex
go build
mv ./runtime ${HOME}/.modus/runtime/v0.16.4/modus_runtime
