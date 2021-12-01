FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /dockergrpc
RUN mkdir -p /dockergrpc/proto
RUN mkdir -p /dockergrpc/proto/proto

WORKDIR /dockergrpc

COPY ./proto/proto/service.pb.go /dockergrpc/proto/proto
COPY ./main.go /dockergrpc

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o dockergrpc .

CMD ./dockergrpc