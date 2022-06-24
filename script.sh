#!/usr/bin/env bash

VERSION=v1.0.2
APPNAME=deploy-hello

# build the project
go build -o ./bin/$APPNAME main.go

# build image
docker build -t birdhk/$APPNAME:$VERSION -f Dockerfile .

# docker login
docker login --username=birdhk --password=zp521314....

# push image to docker hub
docker push birdhk/$APPNAME:$VERSION