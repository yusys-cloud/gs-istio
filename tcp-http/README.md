# RESTful HTTP API communicating with TCP server over TCP network in Golang
In this example we are going to implement a RESTful public HTTP API that communicates with our internal TCP server with a TCP client for storing data. The TCP client code is in our HTTP API. The process is very simple as shown below.

- User sends HTTP request to our public RESTful HTTP API.
- RESTful HTTP API forwards request to internal TCP server with a TCP client.
- TCP server stores the data.
- TCP server returns response to TCP client over TCP network.
- RESTful HTTP API responds to user.

###### User (HTTP request) -> RESTful HTTP API (TCP request) -> TCP Server (TCP response) -> RESTful HTTP API (HTTP response) -> User

## build
``` 
docker build -t='yusyscloud/gs-tcp' . 
docker build -t='yusyscloud/gs-http' . 
```

## tests
``` 
 curl -i -X POST -d '{"username":"username","password":"password"}' "localhost:2002" 
 curl -i -X POST -d '{"username":"username","password":"password"}' "gs-tcp:2002"
```
kubectl logs -f $(kubectl get pod -l app=gs-http-tcp -o jsonpath={.items..metadata.name}) -c gs-tcp