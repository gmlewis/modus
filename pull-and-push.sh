#!/bin/bash -ex
git checkout main
git pull upstream main
git pull upstream main --tags
git push origin main
