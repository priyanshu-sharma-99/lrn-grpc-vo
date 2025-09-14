package main

import (
	"context"
	"io"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{Number: 100}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not calculate: %v", err)
		}
		log.Printf("Prime number: %d", res.Result)
	}
}
