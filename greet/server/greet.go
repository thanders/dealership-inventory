package main

import (
	"context"

	pb "github.com/thanders/dealership-inventory/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.DealerInventory, error) {
	result := getInventoryData()
	return &pb.DealerInventory{
		Inventory: result.Inventory,
	}, nil
}
