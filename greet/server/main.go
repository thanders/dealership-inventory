package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/thanders/dealership-inventory/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
