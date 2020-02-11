package main

import (
	"github.com/k0kubun/pp"
	"github.com/minskylab/tisp"
	sm "github.com/minskylab/tisp/statemanager"
)

func main() {
	state, err := sm.NewStateManager(sm.DefaultPrismaEndpoint, sm.DeafaultPrismaSecret)
	if err != nil {
		panic(err)
	}

	// 5e2ab19de03dd800075f8e48
	// res, err := statemanager.RegisterNewResource(tisp.NewResourceInformation{
	// 	Name: "Bregy Malpartida",
	// 	Cost: tisp.Cost{
	// 		Units: "$/h",
	// 		Value: 10.0,
	// 	},
	// 	MainType: p.ResourceTypeDeveloper,
	// 	Selector: tisp.StrPoint("bregy"),
	// })

	res, err := state.GetResources(tisp.Selector{
		ByIDs: &[]string{"5e2ab19de03dd800075f8e48"},
	})
	if err != nil {
		panic(err)
	}

	pp.Println(res)
}