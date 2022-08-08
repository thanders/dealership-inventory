package main

import (
	"context"
	"log"

	pb "github.com/thanders/dealership-inventory/pkg/proto"
)

func doGetTotalInventory(c pb.HondaInventoryServiceClient) *pb.TotalInventory {
	log.Println("doGetTotalInventory was invoked")
	response, err := c.GetDealerInventory(context.Background(), &pb.GetDealerInventoryRequest{})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	return response
}
