package tg

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func (client *Client) MakeConfigCommand()*cli.Command{



	config := &cli.Command{
		Name:                   "config",
		Usage:                  "configure the tool",
		UsageText:              "tg config --push --file=file.yaml ",
		Description:            "run this command to configure the tools to use credentials provided by tigo pesa in integration stage",
		Before: func(context *cli.Context) error {
			fmt.Printf("checking if the configuration was done before")
			return nil
		},
		After: func(context *cli.Context) error {
			fmt.Printf("printing the config for assurance")
			return nil
		},
		Action: func(context *cli.Context) error {
			fmt.Printf("Write the config into config file")
			return nil
		},
		OnUsageError: func(context *cli.Context, err error, isSubcommand bool) error {
			return nil
		},
		Subcommands:            nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
				Value: "config.yaml",
				Usage: "specify file with config",
			},

			&cli.BoolFlag{
				Name:        "push",
				Usage:       "configuration for push pay",
				Value:       false,
			},
			&cli.BoolFlag{
				Name:        "disburse",
				Usage:       "configuration for disburse",
				Value:       false,
			},
		},

	}

	return config
}

