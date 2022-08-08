package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/thanders/dealership-inventory/di/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)

	}
	defer conn.Close()

	c := pb.NewHondaInventoryServiceClient(conn)

	inventoryList := doGetTotalInventory(c)
	log.Println("Onsite List:")
	createOnsiteTable(inventoryList)
	log.Println("In Transit List:")
	createIntransitTable(inventoryList)

	vehicleNames := doGetVehicleNames(c)
	createVehicleModelsTable(vehicleNames)
}
