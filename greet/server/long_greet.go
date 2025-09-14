package main

import (
	"io"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("LongGreet was invoked!")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: "Hello" + res,
			})
		}

		if err != nil {
			log.Fatalf("could not greet many times: %v", err)
		}

		res += req.FirstName + " "
	}
}
