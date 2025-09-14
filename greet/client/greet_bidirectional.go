package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

func doGreetBidirectional(c pb.GreetServiceClient) {
	log.Println("doGreetBidirectional was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Priyanshu"},
		{FirstName: "Priyanshu2"},
		{FirstName: "Priyanshu3"},
		{FirstName: "Priyanshu4"},
		{FirstName: "Priyanshu5"},
	}

	stream, err := c.BiDirectionGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v", req)
			stream.Send(req)
			time.Sleep(time.Second * 1)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}

			log.Printf("Received: %v", resp)
		}

		close(waitc)
	}()

	<-waitc
	log.Println("doGreetBidirectional was completed")

}
