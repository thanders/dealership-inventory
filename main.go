package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "github.com/thanders/dealership-inventory/proto"
)

type HondaDealerInventory struct {
	IsValidZip bool
	Dealers    []string
	Inventory  []struct {
		DHCost          string
		DealerNumber    string
		ModelYear       string
		ExteriorColor   string
		InteriorColor   string
		ModelGroupName  string
		Transmission    string
		ModelMSRP       string
		NumberOnSite    uint32
		NumberInTransit uint32
	}
	Vehiclemodelseriesname string
	ZipCodeSent            string
}

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

func getRequest(URL string) *http.Response {
	requestURL := fmt.Sprintf(URL)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	return res

}

func getInventoryData() {
	URL := "https://automobiles.honda.com/platform/api/v3/inventoryAndDealers?productDivisionCode=A&modelYear=2022&modelGroup=civic-hatchback&zipCode=85345&maxDealers=5&preferredDealerId=&showOnlineRetailingURL=true"
	result := getRequest(URL)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client: status code: %d\n", result.StatusCode)

	var hondaData HondaDealerInventory
	json.Unmarshal(bytes, &hondaData)

	fmt.Println("Model Name:", hondaData)
	for i := 0; i < len(hondaData.Inventory); i++ {
		fmt.Println("Model Name:")
		fmt.Println("Model Name: " + hondaData.Inventory[i].ModelGroupName)
	}
}

func main() {
	fmt.Println(getCar())
	fmt.Println(getDealership())
	getInventoryData()
}
