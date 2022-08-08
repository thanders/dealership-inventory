package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	pb "github.com/thanders/dealership-inventory/pkg/proto"
)

func createIntransitTable(inventoryList *pb.TotalInventory) {

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', tabwriter.TabIndent)
	fmt.Println("Intransit Inventory")
	fmt.Fprintln(w, "Dealer Number\tModel Name\tModel Year\tIn Transit\tExterior Color\tInterior Color\tTransmission\tMSRP\t")
	for _, v := range inventoryList.InTransitData {
		fmt.Fprint(w, v.DealerNumber, "\t", v.ModelGroupName, "\t", v.ModelYear, "\t", v.NumberInTransit, "\t", v.ExteriorColor, "\t", v.InteriorColor, "\t", v.Transmission, "\t", v.ModelMSRP, "\t\n")
	}
	w.Flush()

}
