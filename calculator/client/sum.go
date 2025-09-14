package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
)

func callSum(c pb.CalculatorServiceClient) {

	log.Println("callSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{FirstNumber: 10, SecondNumber: 20})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	log.Printf("Sum: %d", res.Result)
	log.Println("callSum was completed")
}
