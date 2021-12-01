# Simple Golang gRPC chat with Docker

## Server
Build docker container from project root

```
docker build --tag=dockergrpc .

docker run -it -p 8080:8080 dockergrpc
```

## Client
another terminal window from /client
```
go run main.go -N Someone
```

another terminal window from /client

```
go run main.go -N AnotherOne
```

Write something in any client window, hit enter. 

 