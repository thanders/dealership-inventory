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

func main() {
	fmt.Println(getCar())
}
