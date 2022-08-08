package main

import (
	"context"
	"log"

	pb "github.com/thanders/dealership-inventory/di/proto"
)

func doGetVehicleNames(c pb.HondaInventoryServiceClient) []*pb.HondaVehicleInfo_VehicleModels {
	log.Println("doGetVehicleNames was invoked")
	res, err := c.GetHondaInfo(context.Background(), &pb.GetHondaInfoRequest{})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	return res.Models
}
