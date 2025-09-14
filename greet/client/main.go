package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
)

var add string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(add, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("connected to the remote server")

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)

	log.Println("closing greetclient connection")
	defer conn.Close()
}
