#!/bin/bash
set -x
set -e

cd /go/src
go run main.go
ls -al /var/run/docker.sock || echo "Docker socket not found at /var/run/docker.sock"

