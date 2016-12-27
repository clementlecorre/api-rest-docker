# deploy-docker

Sample Go api For deploy docker container with socket


## Environment variables


| Variabe | Value | Description |
| ------- |:------|:------------|
| `xtoken` |  1234 | xToken API |

## Run


```bash
docker build -t deploy  .
docker run -it -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -e xtoken=0000 deploy
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
