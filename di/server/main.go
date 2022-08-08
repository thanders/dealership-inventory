package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/thanders/dealership-inventory/di/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.HondaInventoryServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)
	s := grpc.NewServer()

	// start service
	pb.RegisterHondaInventoryServiceServer(s, &Server{})

	defer s.Stop()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
