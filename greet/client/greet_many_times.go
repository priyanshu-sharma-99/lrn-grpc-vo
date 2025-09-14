package main

import (
	"context"
	"io"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{FirstName: "Priyanshu"}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet many times: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not greet many times: %v", err)
		}

		log.Printf("Greeting: %s", res.Result)
	}
}
