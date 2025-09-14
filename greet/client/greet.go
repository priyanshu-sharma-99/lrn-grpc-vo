package main

import (
	"context"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Priyanshu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", res.Result)
}
