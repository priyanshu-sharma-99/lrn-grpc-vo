package main

import (
	"context"
	"log"
	"math"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("sqrt was invoked")

	number := in.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument, "number %d is negative", number)
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
