package tg

import (
	"fmt"
	"github.com/techcraftt/tigosdk/aw"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
)

const (
	PushPay RequestType = iota
	Disburse
)

type (
	App struct {
		CLI      *cli.App
		Conf     *Config
		Push     *push.Client
		Disburse *aw.Client
	}
	RequestType int

	Request struct {
		ReferenceID string
		Amount      string
		MSISDN      string
		Remarks     string
	}

	Config struct {
		ReferenceIDPrefix         string
		AppName                   string
		AppVersion                string
		ReleaseDate               string

		DisburseAccountName       string
		DisburseAccountMSISDN     string
		DisburseBrandID           string
		DisbursePIN               string
		DisburseRequestURL        string

		PushUsername              string
		PushPassword              string
		PushPasswordGrantType     string
		PushApiBaseURL            string
		PushGetTokenURL           string
		PushBillerMSISDN          string
		PushBillerCode            string
		PushPayURL                string
		PushReverseTransactionURL string
		PushHealthCheckURL        string
	}
)

func NewApplication()*cli.App{

	author := &cli.Author{
		Name:  "Pius Alfred",
		Email: "pmasengwa@techcraft.co.tz",
	}


	flags := []cli.Flag {
		&cli.BoolFlag{
			Name: "verbose",
			Value: false,
			Usage: "allow verbose output",
		},
	}

	commands := makeCommands()

	a := &cli.App{
		Name:                   "tg",
		Usage:                  "command line tool to execute tigopesa push pay and disbursement requests",
		UsageText:              "tg [config | push | disburse ]",
		Version:                "1.0.0",
		Description:            "use this tool to perform push pay requests or disbursement requests.\nall these requests are through tigopesa.\nmake sure the too is correctly configured via config command",
		Commands:               commands,
		Flags:                  flags,
		EnableBashCompletion:   true,
		Authors:                []*cli.Author{
			author,
		},
		Copyright:              "Pius Alfred (c) 2021. MIT License",


	}

	return a
}

func makeCommands ()[]*cli.Command{
	var cms []*cli.Command

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

	cms = append(cms,config)

	return cms
}