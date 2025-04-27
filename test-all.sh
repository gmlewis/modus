#!/bin/bash

# Exit immediately if any command fails
set -e

# Path to the go.work file
GO_WORK_FILE="go.work"

# Function to run tests and lints in a directory
run_tests_and_lints() {
    local dir="$1"
    echo "Entering directory: $dir"
    pushd "$dir" > /dev/null

    echo "Running 'go test ./...' in $dir"
    if ! go test ./...; then
        echo "Error: 'go test ./...' failed in $dir"
        exit 1
    fi

    echo "Running 'golangci-lint' in $dir"
    if ! golangci-lint run; then
        echo "Error: 'golangci-lint' failed in $dir"
        exit 1
    fi

    popd > /dev/null
    echo "Exited directory: $dir"
}

# Read the go.work file and extract directories under 'use'
inside_use_block=false
while IFS= read -r line; do
    # Trim leading and trailing whitespace
    line=$(echo "$line" | sed -E 's/^[[:space:]]+//;s/[[:space:]]+$//')

    # Check if we're inside the 'use' block
    if [[ $line == "use (" ]]; then
        inside_use_block=true
        continue
    fi

    # Check if we've reached the end of the 'use' block
    if [[ $line == ")" ]]; then
        inside_use_block=false
        continue
    fi

    # If we're inside the 'use' block, process the directory
    if $inside_use_block; then
        # Remove leading './' and any remaining whitespace
        dir=$(echo "$line" | sed -E 's|^\./||;s/[[:space:]]*$//')
        if [[ -n $dir ]]; then
            run_tests_and_lints "$dir"
        fi
    fi
done < "$GO_WORK_FILE"

echo "All tests and lints completed successfully."
