package main

import (
	"os"

	"github.com/minskylab/tisp"
	"github.com/minskylab/tisp/statemanager"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("PRISMA_ENDPOINT", statemanager.DefaultPrismaEndpoint)
	viper.SetDefault("PRISMA_SECRET", statemanager.DeafaultPrismaSecret)

	ctl := &Control{Machine: &tisp.Machine{}}

	state, err := statemanager.NewStateManager(
		viper.GetString("PRISMA_ENDPOINT"),
		viper.GetString("PRISMA_SECRET"),
	)
	if err != nil {
		panic(err)
	}

	ctl.Machine.State = state

	panic(ctl.createCLI().Run(os.Args))
}
