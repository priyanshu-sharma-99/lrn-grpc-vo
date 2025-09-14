package main

import (
	"io"
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Avg(stream grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Println("Avg was invoked")
	var sum int32 = 0
	var count int32 = 0
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			err := stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
			if err != nil {
				log.Println("Error sending message")
			}
			return nil
		}

		count++
		sum += in.Number
		if err != nil {
			log.Fatalf("Failed to receive request: %v", err)
		}
	}
}
