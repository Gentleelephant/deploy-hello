#!/usr/bin/env bash

VERSION=v1.0.0
APPNAME=test-github-action

# build the project
go build -o ./bin/deploy-hello main.go

# build image
docker build -t birdhk/${APPNAME}:${VERSION} .

# docker login
docker login --username=birdhk --password=zp521314....

# push image to docker hub
docker push birdhk/${APPNAME}:${VERSION}