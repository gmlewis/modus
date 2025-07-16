#!/bin/bash -e

source .env.dev.local

echo "Connect to the Neo4J dashboard at: http://localhost:7474/"

docker run \
    --rm \
    --publish=7474:7474 \
    --publish=7687:7687 \
    --env NEO4J_AUTH=${NEO4J_USER}/${NEO4J_PASSWORD} \
    neo4j:5.12.0
