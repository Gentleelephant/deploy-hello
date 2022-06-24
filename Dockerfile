FROM golang:latest

WORKDIR /app

COPY ./bin/deploy-hello /app

LABEL app-name=deploy-hello

ENTRYPOINT ["./deploy-hello"]