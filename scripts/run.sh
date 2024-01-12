#!/bin/sh
echo "running ${1} ${2}"
cd ./services/"${1}" && go run ./cmd/app "${2}"