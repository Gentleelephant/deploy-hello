#!/usr/bin/env bash

# build the project
go build -o ../bin/deploy-hello ../main.go

# build image
cd ../
docker build -t deploy-hello:v1 -f Dockerfile .