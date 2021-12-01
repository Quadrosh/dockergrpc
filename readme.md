# Simple Golang gRPC chat with Docker


To build own "/third_party" folder:
```
protoc --version
go get -u google.golang.org/grpc
go install  github.com/golang/protobuf/protoc-gen-go
```
search "libprotoc" on the disc 

copy "google" folder, which in

protobuf / 3.17.3 (or current) / include / google (on OSX) 

to "/third_party" folder

Run command to build the proto service:
```
protoc --proto_path=proto --proto_path=third_party --go_out=plugins=grpc:proto service.proto
```


## Server
Build docker container from project root, and run it

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

 