FROM golang:1.7

LABEL maintainer "clement@le-corre.eu" \
      go_lib_docker "github.com/yhat/go-docker"\
      go_lib_docker "github.com/gorilla/mux"\
      description "API rest for deploy container easily"

ENV xtoken=1234 \
    docker_version=17.03.1* \
    GOPATH=$GOPATH:/go/api \
    GOBIN=$GOPATH/bin


RUN apt-get update && \
    apt-get -y install \
      apt-transport-https \
      ca-certificates \
      curl \
      software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    add-apt-repository \
         "deb [arch=amd64] https://download.docker.com/linux/debian \
         $(lsb_release -cs) \
         stable" && \
    apt-get update && \
    apt-get -y install docker-ce=$docker_version && \
    rm -rf /var/lib/apt/lists/*

COPY api/ /go/api
WORKDIR /go/api
RUN go get

EXPOSE 8080
CMD ["go", "run", "*.go"]
