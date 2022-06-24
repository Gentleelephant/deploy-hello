#!/usr/bin/env bash

# build the project
go build -o ./bin/deploy-hello main.go

# build image
docker build -t birdhk/deploy-hello:v1 -f Dockerfile .

# docker login
docker login --username=birdhk --password=zp521314....

# push image to docker hub
docker push birdhk/deploy-hello:v1