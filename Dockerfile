FROM golang:1.7

MAINTAINER cl3m3nt <clement@le-corre.eu>

RUN apt-get update \
  && apt-get -y install wget vim \
  && wget -qO- https://get.docker.com/ | sh

ENV xtoken 1234
ENV GOBIN=$GOPATH/bin
COPY /api /go
WORKDIR /go
EXPOSE 8080

RUN go get 
CMD go run *.go
