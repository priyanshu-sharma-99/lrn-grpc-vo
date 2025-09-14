package main

import (
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	log.Printf("GreetManyTimes was invoked with: %v \n", in)

	for i := 0; i < 10; i++ {
		stream.Send(&pb.GreetResponse{
			Result: "Hello " + in.FirstName,
		})
	}
	return nil
}
