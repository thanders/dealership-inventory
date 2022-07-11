package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	pb "github.com/thanders/dealership-inventory/greet/proto"
)

func getHondaModels() *pb.HondaVehicleInfo {
	URL := "https://automobiles.honda.com/platform/api/v1/search-inventory/configuration"
	result := getRequest(URL)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Request status code: ", result.StatusCode)

	var data *pb.HondaVehicleInfo
	json.Unmarshal(bytes, &data)

	hondaModels := &pb.HondaVehicleInfo{
		Models: data.Models,
	}

	return hondaModels
}
