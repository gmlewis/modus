#!/bin/bash -ex
go test ./...
go run ./tools/modus-moonbit-build -test tools/modus-moonbit-build/end-to-end-tests.json
