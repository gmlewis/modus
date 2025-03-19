#!/bin/bash -ex

go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2

export POSTGRESQL_URL='postgresql://postgres:postgres@localhost:5433/my-runtime-db?sslmode=disable'
migrate -database ${POSTGRESQL_URL} -path ../../../../runtime/db/migrations up
