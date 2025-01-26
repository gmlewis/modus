#!/bin/bash -ex

original_dir=$(pwd)
for dir in */; do
    # Remove trailing slash
    dir=${dir%/}
    cd "$dir"
    if [ -x "./update.sh" ]; then
        ./update.sh
    fi
    cd "$original_dir"
done
