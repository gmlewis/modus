#!/bin/bash -ex

# First, open up a separate terminal window, and run this script.
pushd ../../../../runtime/tools/local
docker compose up

# Then, in another terminal window, migrate the database by running:
# ./migrate-database-once.sh
