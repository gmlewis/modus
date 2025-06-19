#!/bin/bash -x

original_dir=$(pwd)
for dir in */; do
    # Remove trailing slash
    dir=${dir%/}
    cd "$dir"
    if [ -x "./build.sh" ]; then
        ./build.sh
    fi
    cd "$original_dir"
done
