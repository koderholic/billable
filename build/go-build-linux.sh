#!/bin/bash
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o billable.elf -ldflags "-s -w" && upx billable.elf