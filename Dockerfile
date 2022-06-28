FROM alpine:latest

WORKDIR /app

COPY ./deploy-hello /app

LABEL app-name=deploy-hello

ENTRYPOINT ["./deploy-hello"]