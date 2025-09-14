package main

import (
	"context"
	"log"
	"time"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.GreetServiceClient) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // will cancel the context when the function returns

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{FirstName: "Priyanshu"})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
			} else {
				log.Fatalf("Unexpected grpc error: %v", err)
			}
		} else {
			log.Printf("A non-grpc error occurred: %v", err)
		}
		return
	}

	log.Printf("GreetWithDeadline response: %v", res.Result)
}
