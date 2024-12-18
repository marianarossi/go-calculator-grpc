package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ibilalkayy/Small-Projects/gRPC/calculator/proto" //biblioteca de calculadora do GO
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.Response, error) {
	return &pb.Response{Response: req.Num1 + req.Num2}, nil
}

func (s *server) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.Response, error) {
	return &pb.Response{Response: req.Num1 - req.Num2}, nil
}

func (s *server) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.Response, error) {
	return &pb.Response{Response: req.Num1 * req.Num2}, nil
}

func (s *server) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.Response, error) {
	return &pb.Response{Response: req.Num1 / req.Num2}, nil
}

func main() {
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Printf("Server is running on port %s...\n", port)

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &server{})
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
