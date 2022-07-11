package main

import (
	"context"
	"log"

	pb "github.com/thanders/dealership-inventory/greet/proto"
)

func doGreet(c pb.HondaInventoryServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.GetAllData(context.Background(), &pb.GetAllDataRequest{
		FirstName: "Tom",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res)

	log.Printf("Model: %s\n", res.Inventory[0].ModelGroupName)
}

func doGetTotalInventory(c pb.HondaInventoryServiceClient) {
	log.Println("doGetTotalInventory was invoked")
	res, err := c.GetDealerInventory(context.Background(), &pb.GetDealerInventoryRequest{})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Println("GET TOTAL INVENTORY RESULT --- :", res)
	/*
		for _, r := range res.InTransitData {
			log.Println("Number in Transit: ", r.NumberInTransit)
		}
	*/
}
