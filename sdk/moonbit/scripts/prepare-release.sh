#!/usr/bin/env bash
# Prepares templates for release.

set -euo pipefail
trap "cd \"${PWD}\"" EXIT
cd "$(dirname "$0")"

# get the version number from the command line argument
if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <version>"
  exit 1
fi
version=$1

version=${version#"v"}
echo "Preparing release for version ${version}"
cd ..

# Update the version in the sdk.json file
jq --arg ver "${version}" '.sdk.version = $ver' sdk.json > tmp.json && mv tmp.json sdk.json

# Update the version of the sdk used in the templates
cd templates
for template in *; do
  if [ -d "${template}" ]; then
    cd "${template}"

    # TODO

    cd ..
  fi
done
cd ..

# Create a tarball of the templates
tar -czvf templates_moonbit_v${version}.tar.gz templates
