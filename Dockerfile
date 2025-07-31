#FROM golang:latest AS builder
#
#COPY . /src
#WORKDIR /src
#
#RUN GOPROXY=https://goproxy.cn make build
#
FROM debian:latest

COPY ./bin/linux-amd64 /app/bin
COPY ./cmd/server/app.yml /app/conf/app.yml

WORKDIR /app

EXPOSE 9888
EXPOSE 19888
VOLUME /app/conf

CMD ["/app/bin/server", "-Dapp.conf=/app/conf/app.yml"]
