package main

import (
	"log"
	"os"

	"github.com/ezkl/adr/pkg/helpers"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "adr"
	app.Usage = "Work with Architecture Decision Records (ADRs)"
	app.Version = "0.1.0"

	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{}
	app.Commands = []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"c"},
			Usage:   "Create a new ADR",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) error {
				cfg := helpers.GetConfig()
				cfg.CurrentAdr++

				err := helpers.NewAdr(cfg, c.Args().Slice())

				if err != nil {
					return err
				}

				err = helpers.UpdateConfig(cfg)

				if err != nil {
					return err
				}

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
