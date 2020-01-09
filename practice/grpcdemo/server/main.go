package main

import (
	"context"
	"fmt"
	"github.com/chisty/grpcdemo-go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type calculationServer struct{}

func main() {
	fmt.Println("Hello GRPC")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Unable to listen on port 5001")
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCalculationServiceServer(grpcServer, &calculationServer{})
	reflection.Register(grpcServer)
	fmt.Println("Starting Server...")
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *calculationServer) Add(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()
	fmt.Printf("Add Request with value %d and %d\n", a, b)
	return &proto.Response{Result: a + b}, nil
}

func (s *calculationServer) Multiply(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	a, b := req.GetA(), req.GetB()
	fmt.Printf("Multiply Request with value %d and %d\n", a, b)
	return &proto.Response{Result: a * b}, nil
}
