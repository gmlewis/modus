#!/bin/bash -ex

docker run \
    --rm \
    -it \
    -p "8080:8080" \
    -p "9080:9080" \
    -v ~/dgraph:/dgraph \
    "dgraph/standalone:v24.1.0"

#    -p "8000:8000" \
