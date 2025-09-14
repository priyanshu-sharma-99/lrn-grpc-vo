package main

import (
	"log"
	"net"

	pb "github.com/priyanshu-sharma-99/lrn-grpc-v0/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var add string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", lis.Addr())
	tls := true

	opts := []grpc.ServerOption{}

	if tls {

		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("failed to generate credentials %v", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Println("greet server stopped")
}
