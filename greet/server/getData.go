package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "github.com/thanders/dealership-inventory/greet/proto"
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

type DealerData struct {
	ASARanking string
	Address    string
}

type DealerInventory struct {
	IsValidZip             bool
	Dealers                []DealerData
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

func createInventoryPB(inventoryArray []*pb.DealerInventory_InventoryData) []*pb.DealerInventory_InventoryData {
	di := make([]*pb.DealerInventory_InventoryData, len(inventoryArray))
	for i, r := range inventoryArray {
		di[i] = &pb.DealerInventory_InventoryData{
			DHCost:          r.DHCost,
			DealerNumber:    r.DealerNumber,
			ModelYear:       r.ModelYear,
			ExteriorColor:   r.ExteriorColor,
			InteriorColor:   r.InteriorColor,
			ModelGroupName:  r.ModelGroupName,
			Transmission:    r.Transmission,
			ModelMSRP:       r.ModelMSRP,
			NumberOnSite:    r.NumberOnSite,
			NumberInTransit: r.NumberInTransit,
		}
	}
	return di
}

func createDealersPB(array []*pb.DealerInventory_DealerData) []*pb.DealerInventory_DealerData {
	dealers := make([]*pb.DealerInventory_DealerData, len(array))
	for i, r := range array {
		dealers[i] = &pb.DealerInventory_DealerData{
			ASARanking:           r.ASARanking,
			Address:              r.Address,
			City:                 r.City,
			DealerDistanceMiles:  r.DealerDistanceMiles,
			DealerNumber:         r.DealerNumber,
			DrivingDistanceMiles: r.DrivingDistanceMiles,
			InternetManager:      r.InternetManager,
			IsServiceCenter:      r.IsServiceCenter,
			Latitude:             r.Latitude,
			Longitude:            r.Longitude,
			Name:                 r.Name,
			PartsHours:           r.PartsHours,
			PartsPhone:           r.PartsPhone,
			Phone:                r.Phone,
			SalesHours:           r.SalesHours,
			ServiceHours:         r.ServiceHours,
			ServicePhone:         r.ServicePhone,
			State:                r.State,
			WebAddress:           r.WebAddress,
			ZipCode:              r.ZipCode,
		}
	}
	return dealers
}

func getInventoryData() *pb.DealerInventory {
	URL := "https://automobiles.honda.com/platform/api/v3/inventoryAndDealers?productDivisionCode=A&modelYear=2022&modelGroup=civic-hatchback&zipCode=85705&maxDealers=5&preferredDealerId=&showOnlineRetailingURL=true"
	result := getRequest(URL)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client: status code: %d\n", result.StatusCode)

	var data *pb.DealerInventory
	json.Unmarshal(bytes, &data)

	inventorypb := createInventoryPB(data.Inventory)
	dealerspb := createDealersPB(data.Dealers)
	cars := &pb.DealerInventory{
		IsValidZip:             data.IsValidZip,
		Dealers:                dealerspb,
		Inventory:              inventorypb,
		Vehiclemodelseriesname: data.Vehiclemodelseriesname,
		ZipCodeSent:            data.ZipCodeSent,
	}

	return cars
}

func createOnSitePB(array []*pb.DealerInventory_InventoryData) []*pb.TotalInventory_Onsite {
	dealers := make([]*pb.TotalInventory_Onsite, len(array))
	for i, r := range array {
		if r.NumberOnSite > 0 {
			dealers[i] = &pb.TotalInventory_Onsite{
				DHCost:          r.DHCost,
				DealerNumber:    r.DealerNumber,
				ModelYear:       r.ModelYear,
				ExteriorColor:   r.ExteriorColor,
				InteriorColor:   r.InteriorColor,
				ModelGroupName:  r.ModelGroupName,
				Transmission:    r.Transmission,
				ModelMSRP:       r.ModelMSRP,
				NumberOnSite:    r.NumberOnSite,
				NumberInTransit: r.NumberInTransit,
			}
		}
	}
	return dealers
}

func createInTransitPB(array []*pb.DealerInventory_InventoryData) []*pb.TotalInventory_InTransit {

	var inTransitLength uint32 = 0
	for i := 0; i < len(array); i++ {
		if array[i].NumberInTransit > 0 {
			inTransitLength += array[i].NumberInTransit
		}
	}
	fmt.Println("LENGTH RESULT --- ", inTransitLength)
	fmt.Println("LENGTH Array --- ", len(array))
	dealers := make([]*pb.TotalInventory_InTransit, inTransitLength)
	for i, r := range array {
		fmt.Println("loop --- ", i)
		if r.NumberInTransit > 0 {
			fmt.Println("bigger than zero --- ", i, r.NumberInTransit)
			dealers[i] = &pb.TotalInventory_InTransit{
				DHCost:          r.DHCost,
				DealerNumber:    r.DealerNumber,
				ModelYear:       r.ModelYear,
				ExteriorColor:   r.ExteriorColor,
				InteriorColor:   r.InteriorColor,
				ModelGroupName:  r.ModelGroupName,
				Transmission:    r.Transmission,
				ModelMSRP:       r.ModelMSRP,
				NumberOnSite:    r.NumberOnSite,
				NumberInTransit: r.NumberInTransit,
			}
		}
	}
	return dealers
}

func getTotalInventory() *pb.TotalInventory {
	URL := "https://automobiles.honda.com/platform/api/v3/inventoryAndDealers?productDivisionCode=A&modelYear=2022&modelGroup=civic-hatchback&zipCode=85705&maxDealers=5&preferredDealerId=&showOnlineRetailingURL=true"
	result := getRequest(URL)

	bytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client: status code: %d\n", result.StatusCode)

	var data *pb.DealerInventory
	json.Unmarshal(bytes, &data)

	var totalInTransit uint32 = 0
	var totalOnSite uint32 = 0
	for _, r := range data.Inventory {
		totalInTransit = totalInTransit + r.NumberInTransit
		totalOnSite = totalOnSite + r.NumberOnSite
	}

	onSitePB := createOnSitePB(data.Inventory)
	inTransitPB := createInTransitPB(data.Inventory)

	fmt.Println("LENGTH inTransitPB --- ", len(inTransitPB))

	totalInventory := &pb.TotalInventory{
		ModelGroupName: data.Vehiclemodelseriesname,
		TotalOnsite:    totalOnSite,
		TotalIncoming:  totalInTransit,
		ZipCode:        data.ZipCodeSent,
		OnSiteData:     onSitePB,
		InTransitData:  inTransitPB,
	}

	return totalInventory
}
