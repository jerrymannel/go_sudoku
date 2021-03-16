#!/bin/bash
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o goSudoku .