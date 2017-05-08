# api-rest-docker

[![Docker Automated buil](https://img.shields.io/docker/automated/cl3m3nt/api-rest-docker.svg?style=flat-square)](https://hub.docker.com/r/cl3m3nt/api-rest-docker/)
[![Docker Build Statu](https://img.shields.io/docker/build/cl3m3nt/api-rest-docker.svg?style=flat-square)](https://hub.docker.com/r/cl3m3nt/api-rest-docker/)
[![Docker Pulls](https://img.shields.io/docker/pulls/cl3m3nt/api-rest-docker.svg?style=flat-square)](https://hub.docker.com/r/cl3m3nt/api-rest-docker/)

Deploy docker container with Go rest api

## Library used

* https://github.com/yhat/go-docker
* https://github.com/gorilla/mux

## Environment variables


| Variabe | Value | Description |
| ------- |:------|:------------|
| `xtoken` |  1234 | xToken API |

## Run


```bash
docker build -t cl3m3nt/api-rest-docker  .
docker run -it -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -e xtoken=0000 cl3m3nt/api-rest-docker
```

## Deploy with API


```bash
curl -X POST --header 'Content-Type: application/json' --header 'X-AUTH-TOKEN: 0000' -d '{
    "Image": "httpd:alpine",
    "ExposedPorts": {
        "80/tcp" : {}
    },
  "HostConfig":{
    "PortBindings": {
              "80/tcp": [
                  {
                      "HostIp": "0.0.0.0",
                      "HostPort": "80"
                  }
              ]
          }
  },
  "Hostname": "coucou",
  "Env": [
                "test=coucou"
    ]
}' http://localhost:8080/deploy
```

```bash
curl -X POST --header 'Content-Type: application/json' --header 'X-AUTH-TOKEN: 1234' -d '{
    "Image": "nginx",
    "Env": [
                "VIRTUAL_NETWORK=nginx-proxy",
                "LETSENCRYPT_HOST=test.mydomain.com",
                "LETSENCRYPT_EMAIL=clement@mydomain.com",
                "VIRTUAL_HOST=test.mydomain.com",
                "VIRTUAL_PORT=80"
            ],
  "HostConfig":{
    "PortBindings": {}
  }
  "Hostname": "test"
}' http://localhost:8080/deploy
```
