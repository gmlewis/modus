#!/bin/bash -ex
go test ./...
go run . -test end-to-end-tests.json
