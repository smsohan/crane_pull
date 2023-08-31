#!/bin/bash
set -x
set -e

cd /go/src
go run main.go
ls -al /var/run/docker.sock

