FROM golang:1.7

MAINTAINER cl3m3nt <clement@le-corre.eu>

RUN apt-get update \
  && apt-get -y install wget vim \
  && wget -qO- https://experimental.docker.com/ | sh
ENV xtoken 1234
#COPY /api /go
#RUN export GOBIN=$GOPATH/bin && go get

EXPOSE 8080
