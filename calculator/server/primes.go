package main

import (
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
	"google.golang.org/grpc"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	log.Printf("Primes was invoked with: %v \n", in)

	number := in.Number
	divisor := int64(2)
	for number > 1 {
		if number%divisor == 0 {
			err := stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})
			if err != nil {
				log.Println("Error sending message")
				return err
			}
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
