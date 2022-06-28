#!/usr/bin/bash

VERSION=v1.0.0
APPNAME=deploy-hello

# build the project
echo "Building the project..."
go build ./main.go
ls ./

if [ ! -e ./deploy-hello ]; then
    echo "Error: deploy-hello not found"
    exit 1
fi
# build image
docker build -t birdhk/${APPNAME}:${VERSION} .

# docker login
docker login --username=birdhk --password=zp521314....

# push image to docker hub
docker push birdhk/${APPNAME}:${VERSION}