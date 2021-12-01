package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"

	"github.com/quadrosh/dockergrpc/proto/proto"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

// Connection ...
type Connection struct {
	stream proto.Broadcast_CreateStreamServer
	id     string
	active bool
	error  chan error
}

// Server ...
type Server struct {
	Connection []*Connection
}

// CreateStream ...
func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream: stream,
		id:     pconn.User.Id,
		active: true,
		error:  make(chan error),
	}
	s.Connection = append(s.Connection, conn)
	return <-conn.error
}

// BroadcastMessage ...
func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{} // counts the goroutines in progress, wait till it finished and decrements
	done := make(chan int)

	for _, conn := range s.Connection {
		wait.Add(1)

		go func(msg *proto.Message, conn *Connection) {
			defer wait.Done() // decrement waitgroup
			if conn.active {
				err := conn.stream.Send(msg)
				grpcLog.Info("sending message to: ", conn.stream)
				if err != nil {
					grpcLog.Errorf("Error with Stream: %s - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)
	}

	go func() { // goroutine
		wait.Wait() // wait till an other goroutines will exit
		close(done) // than execute close command to done channel
	}()

	<-done // block the return from function untill done retuns an item
	return &proto.Close{}, nil
}
func main() {
	var connections []*Connection

	server := &Server{connections}

	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error creating the server %v", err)
	}

	grpcLog.Info("Starting server at port :8080")

	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
}
