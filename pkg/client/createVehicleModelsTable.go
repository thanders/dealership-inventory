package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	pb "github.com/thanders/dealership-inventory/pkg/proto"
)

func createVehicleModelsTable(vehicleModels []*pb.HondaVehicleInfo_VehicleModels) {

	w := tabwriter.NewWriter(os.Stdout, 10, 0, 4, ' ', tabwriter.TabIndent)
	fmt.Println("Vehicle Models")
	fmt.Fprintln(w, "ModelName\tModelYear\tMSRP\tMPG\t")
	for _, v := range vehicleModels {
		fmt.Fprint(w, v.ModelName, "\t", v.ModelYear, "\t", v.Msrp, "\t", v.Mpg, "\t\n")
	}
	w.Flush()

}
