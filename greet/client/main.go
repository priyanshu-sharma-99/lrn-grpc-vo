package main

import (
	"log"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var add string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading cert credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(add, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	log.Println("connected to the remote server")

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	doGreetManyTimes(c)
	doLongGreet(c)
	doGreetBidirectional(c)
	doGreetWithDeadline(c)

	log.Println("closing greetclient connection")
	defer conn.Close()
}
