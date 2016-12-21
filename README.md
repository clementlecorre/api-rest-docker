# deploy-docker

Sample Go api For deploy docker container with socket

## Run

```
docker build -t deploy  .
docker run -it -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v $PWD/api:/go deploy bash
export GOBIN=$GOPATH/bin && go get
export GOBIN=$GOPATH/bin && go run *.go
```
