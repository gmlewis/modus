#!/bin/bash -ex

MOONBIT_SDK_VERSION=v0.16.5
go build -ldflags "-checklinkname=0"
mv runtime ${HOME}/.modus/runtime/${MOONBIT_SDK_VERSION}/modus_runtime

pushd languages/moonbit/testdata
./build.sh
popd
