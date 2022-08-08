package main

import (
	"context"
	"log"

	pb "github.com/thanders/dealership-inventory/pkg/proto"
)

func (s *Server) GetAllData(ctx context.Context, in *pb.GetAllDataRequest) (*pb.DealerInventory, error) {
	log.Printf("Greet was invoked with %v\n", in)
	result := getInventoryData()
	return result, nil
}

func (s *Server) GetDealerInventory(ctx context.Context, in *pb.GetDealerInventoryRequest) (*pb.TotalInventory, error) {
	result := getTotalInventory()
	return result, nil
}

func (s *Server) GetHondaInfo(ctx context.Context, in *pb.GetHondaInfoRequest) (*pb.HondaVehicleInfo, error) {
	result := getHondaModels()
	return result, nil
}
