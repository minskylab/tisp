package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)
func createCLI() *cli.App {
	app := &cli.App{
		Name: "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
	}

	return app
}
