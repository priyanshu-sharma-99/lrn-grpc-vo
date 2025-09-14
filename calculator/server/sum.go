package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Sum was invoked")
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
