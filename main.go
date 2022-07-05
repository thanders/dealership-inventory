package main

import (
	"fmt"

	pb "github.com/thanders/dealership-inventory/proto"
)

func getCar() *pb.Car {
	return &pb.Car{
		Id:           28,
		IsManual:     true,
		Manufacturer: "Honda",
		Model:        "Civic",
		SampleLists:  []int32{1, 2, 3, 4, 5, 6},
	}
}

func getDealership() *pb.Dealership {
	civic := &pb.Car{
		Id:           28,
		IsManual:     true,
		Manufacturer: "Honda",
		Model:        "Civic",
		SampleLists:  []int32{1, 2, 3, 4, 5, 6},
	}

	accord := &pb.Car{
		Id:           14,
		IsManual:     true,
		Manufacturer: "Honda",
		Model:        "Accord",
		SampleLists:  []int32{1, 2, 3, 4, 5, 6},
	}

	return &pb.Dealership{
		Name:    "Test Honda",
		Address: "Test address",
		Inventory: []*pb.Car{
			civic,
			accord,
		},
		Sedans: &pb.Vehicles_Sedan{
			Civic:  []*pb.Car{civic},
			Accord: []*pb.Car{accord},
		}, // Sedans are a Vehicle object with different models, each model corresponds to an array of cars
	}
}

func funcGetInventory() {

}

func main() {
	fmt.Println(getCar())
	fmt.Println(getDealership())
}
