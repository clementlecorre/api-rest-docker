# deploy-docker

Sample Go api For deploy docker container with socket

## Run

```
docker build -t deploy  .
docker run -it -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -v $PWD/api:/go deploy bash
export GOBIN=$GOPATH/bin && go get
export GOBIN=$GOPATH/bin && go run *.go
```

## TEST

```
curl -X POST --header 'Content-Type: application/json' --header 'X-AUTH-TOKEN: 1234' -d '{
	"Image": "httpd:alpine",
	"PortSpecs": [
		"80/tcp"
	],
	"ExposedPorts": [
		"80/tcp"
	],
  "HostConfig":{
      "PortBindings": {
         "80/tcp":[{"HostPort":"80"}]
       }
  },
	"Hostname": "coucou"
}' http://localhost:8080/deploy
```
