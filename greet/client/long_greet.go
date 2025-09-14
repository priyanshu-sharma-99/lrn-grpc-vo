package main

import (
	"context"
	"log"
	"time"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("DoLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Priyanshu"},
		{FirstName: "Mohit"},
		{FirstName: "Ashish"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("could not greet many times: %v", err)
	}

	for _, req := range reqs {
		log.Println("sending request")
		stream.Send(req)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not greet many times: %v", err)
	}

	log.Printf("GreetManyTimes response: %v", res)
}
