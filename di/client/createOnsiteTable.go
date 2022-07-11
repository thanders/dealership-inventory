package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	pb "github.com/thanders/dealership-inventory/di/proto"
)

func createOnsiteTable(inventoryList *pb.TotalInventory) {

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', tabwriter.TabIndent)
	fmt.Println("ONSITE Inventory")
	fmt.Fprintln(w, "ModelName\tModelYear\tOnsite\tExterior Color\tInterior Color\tTransmission\tMSRP\t")
	for _, v := range inventoryList.OnSiteData {
		fmt.Fprint(w, v.ModelGroupName, "\t", v.ModelYear, "\t", v.NumberOnSite, "\t", v.ExteriorColor, "\t", v.InteriorColor, "\t", v.Transmission, "\t", v.ModelMSRP, "\t\n")
	}
	w.Flush()

}
