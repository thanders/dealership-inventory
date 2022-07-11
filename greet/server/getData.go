package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type InventoryData struct {
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

type DealerInventory struct {
	IsValidZip             bool
	Dealers                []string
	Inventory              []InventoryData
	Vehiclemodelseriesname string
	ZipCodeSent            string
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

func getInventoryData() DealerInventory {
	URL := "https://automobiles.honda.com/platform/api/v3/inventoryAndDealers?productDivisionCode=A&modelYear=2022&modelGroup=civic-hatchback&zipCode=85345&maxDealers=5&preferredDealerId=&showOnlineRetailingURL=true"
	result := getRequest(URL)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client: status code: %d\n", result.StatusCode)

	var data DealerInventory
	json.Unmarshal(bytes, &data)
	return data
	/*
		fmt.Println("Model Name:", data)
		for i := 0; i < len(data.Inventory); i++ {
			fmt.Println("Model Name:")
			fmt.Println("Model Name: " + data.Inventory[i].ModelGroupName)
		}
	*/
}
