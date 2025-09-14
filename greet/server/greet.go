package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", in.GetFirstName())
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
