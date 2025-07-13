#!/bin/bash -ex
NO_COLOR=1 go test ./... 2>&1 | grep 'FAIL: '
