#!/bin/bash -ex
npm install
npm run build
./bin/modus.js -h sdk
