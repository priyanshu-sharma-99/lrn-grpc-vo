package main

import (
	"io"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) BiDirectionGreet(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("BiDirectionGreet was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error reading client streams: %v", err)
		}

		res := "Hello " + req.FirstName
		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("Error writing client streams: %v", err)
		}
	}
}
