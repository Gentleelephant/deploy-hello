FROM alpine:latest

WORKDIR /app

COPY ./deploy-hello /app

COPY ./start.sh  /app

LABEL app-name=deploy-hello

ENTRYPOINT ["sh","./start.sh"]