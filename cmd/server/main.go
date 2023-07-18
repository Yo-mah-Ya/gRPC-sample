package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	hellopb "sample/pkg/grpc"

	"google.golang.org/grpc"
)

type myServer struct {
	hellopb.UnimplementedGreetingServiceServer
}

func NewMyServer() *myServer {
	return &myServer{}
}

func (s *myServer) Hello(ctx context.Context, req *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	return &hellopb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.GetName()),
		Age:     req.GetAge(),
	}, nil
}

func main() {
	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	// register GreetingService into gRPC server
	hellopb.RegisterGreetingServiceServer(grpcServer, NewMyServer())

	go func() {
		log.Printf("start gRPC server port: %v", port)
		grpcServer.Serve(listener)
	}()

	// When ctrl + C is typed, execute Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	grpcServer.GracefulStop()
}
