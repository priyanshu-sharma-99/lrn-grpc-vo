package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/calculator/proto"
)

var add string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("connected to the remote server")

	c := pb.NewCalculatorServiceClient(conn)

	callSum(c)
	doPrimes(c)

	log.Println("closing calculatorClient connection")
	defer conn.Close()
}
