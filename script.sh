#!/usr/bin/env bash

# build the project
go build -o ./bin/deploy-hello main.go

# build image
docker build -t deploy-hello:v1 -f Dockerfile .

# push image to docker hub
docker push deploy-hello:v1