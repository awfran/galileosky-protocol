#!/usr/bin/env bash

OUTPUT_PATH=./bin

mkdir -p $OUTPUT_PATH
go build -o $OUTPUT_PATH/receiver ../cmd/receiver