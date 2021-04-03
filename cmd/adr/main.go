package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/fatih/color"
	"github.com/ezkl/adr/pkg/helpers"
)

func main() {
	app := cli.NewApp()
	app.Name = "adr"
	app.Usage = "Work with Architecture Decision Records (ADRs)"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"c"},
			Usage:   "Create a new ADR",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) error {
				currentConfig := helpers.GetConfig()
				currentConfig.CurrentAdr++
				helpers.UpdateConfig(currentConfig)
				helpers.NewAdr(currentConfig, c.Args())
				return nil
			},
		},

		{
			Name:        "init",
			Aliases:     []string{"i"},
			Usage:       "Initializes the ADR configurations",
			UsageText:   "adr init /home/user/adrs",
			Description: "Initializes the ADR configuration with an optional ADR base directory\n This is a a prerequisite to running any other adr sub-command",
			Action: func(c *cli.Context) error {
				initDir := c.Args().First()

				if initDir == "" {
					initDir = helpers.AdrDefaultBaseFolder
				}

				color.Green("Initializing ADR base at " + initDir)
				helpers.InitBaseDir(initDir)
				helpers.InitConfig(initDir)
				helpers.InitTemplate()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
