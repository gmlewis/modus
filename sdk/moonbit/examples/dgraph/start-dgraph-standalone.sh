#!/bin/bash -ex

docker run \
    --rm \
    -it \
    -p "8080:8080" \
    -p "9080:9080" \
    -p "8000:8000" \
    -v ~/dgraph:/dgraph \
    "dgraph/standalone:v24.0.5"
