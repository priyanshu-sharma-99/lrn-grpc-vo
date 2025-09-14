package main

import (
	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
	"google.golang.org/grpc"

	"log"
	"net"
)

var add string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("calculator server stopped")
}
