#!/usr/bin/bash

mkdir -p build
GOOS=linux GOARCH=arm64 go build -o build/jarate cmd/jarate/main.go
