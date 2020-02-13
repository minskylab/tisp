package main

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func (ctl *Control) createCLI() *cli.App {
	var filepath string
	app := &cli.App{
		Name: "tispctl",
		Usage: "tispctl is a cli tool to interact with your tisp instance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:                   "apply",
				Usage:                  "apply a yaml file configurator",
				Description:            "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name: "file",
						Aliases: []string{"f"},
						Value: "",
						Destination: &filepath,
					},
				},
				Action: func(c *cli.Context) error {
					if c.NumFlags() < 1 {
						return errors.New("file flag not found")
					}

					return ctl.apply(filepath)
				},
			},
		},
	}

	return app
}
