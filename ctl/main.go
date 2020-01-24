package main

import (
	"fmt"

	// "github.com/minskylab/tisp"
	"github.com/minskylab/tisp/db/prisma"
	// p "github.com/minskylab/tisp/dbclient"
)

func main() {
	db, err := prisma.NewDB()
	if err != nil {
		panic(err)
	}

	// 5e2ab19de03dd800075f8e48
	// res, err := db.RegisterNewResource(tisp.NewResourceInformation{
	// 	Name: "Bregy Malpartida",
	// 	Cost: tisp.Cost{
	// 		Units: "$/h",
	// 		Value: 10.0,
	// 	},
	// 	MainType: p.ResourceTypeDeveloper,
	// 	Selector: tisp.StrPoint("bregy"),
	// })
	res, err := db.GetResource("5e2ab19de03dd800075f8e48")
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
