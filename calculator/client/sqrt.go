package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	req := &pb.SqrtRequest{Number: n}

	_, err := c.Sqrt(context.Background(), req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("Error code: %v", e.Code())
			log.Printf("Error message: %v", e.Message())

			if e.Code() == codes.InvalidArgument {
				log.Println("Our Mistake, we sent a negative number")
			}
		} else {
			log.Fatalf("A non-grpc error: %v", err)
		}
	} else {
		log.Println("doSqrt was completed")
	}
}
