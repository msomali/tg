package tg

import (
	"github.com/techcraftt/tigosdk/aw"
	"github.com/techcraftt/tigosdk/pkg/tigo"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
)

const (
	PushPay RequestType = iota
	Disburse
)



type (

	Option func(app *App)

	App struct {
		cli *cli.App
		client *Client
	}
	Client struct {
		push     *push.Client
		disburse *aw.Client
	}
	RequestType int

	Request struct {
		ReferenceID string
		Amount      int
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

func Make(config *Config, opts...Option)*App{

	baseClient := tigo.NewBaseClient()

	disburseConfig := &aw.Config{
		AccountName:   config.DisburseAccountName,
		AccountMSISDN: config.DisburseAccountMSISDN,
		BrandID:       config.DisburseBrandID,
		PIN:           config.DisbursePIN,
		RequestURL:    config.DisburseRequestURL,
	}
	disburseClient := &aw.Client{
		Config:     disburseConfig,
		BaseClient: baseClient,
	}

	pushConfig := &push.Config{
		Username:              config.PushUsername,
		Password:              config.PushPassword,
		PasswordGrantType:     config.PushPasswordGrantType,
		ApiBaseURL:            config.PushApiBaseURL,
		GetTokenURL:           config.PushGetTokenURL,
		BillerMSISDN:          config.PushBillerMSISDN,
		BillerCode:            config.PushBillerCode,
		PushPayURL:            config.PushPayURL,
		ReverseTransactionURL: config.PushReverseTransactionURL,
		HealthCheckURL:        config.PushHealthCheckURL,
	}

	pushClient := &push.Client{
		Config:          pushConfig,
		BaseClient:      baseClient,
		CallbackHandler: nil,
	}

	client := &Client{
		push:     pushClient,
		disburse: disburseClient,
	}
	app := &App{
		cli:    app(client),
		client: client,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func app(client *Client)*cli.App{

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

	var commands []*cli.Command

	commands = append(
		commands,
		client.MakePushCommand(),
		client.MakeDisburseCommand(),
	)


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


func (app *App) Run(args []string)error{
	return app.cli.Run(args)
}