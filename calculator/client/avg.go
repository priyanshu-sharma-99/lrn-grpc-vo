package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("do_avg is invoked")
	reqs := []*pb.AvgRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
		{Number: 5},
	}

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}

	for _, req := range reqs {
		log.Printf("sending request: %v", req)
		stream.SendMsg(req)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not calculate: %v", err)
	}
	log.Printf("Average: %f", resp.Result)
}
