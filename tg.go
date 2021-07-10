/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package tg

import (
	"errors"
	"github.com/techcraftt/tigosdk/aw"
	"github.com/techcraftt/tigosdk/pkg/tigo"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
)

const (
	PushPay RequestType = iota
	Disburse
)

var (
	ErrNilEmptyConfigVar   = errors.New("nil or empty configuration variable")
	defaultResponsePrinter = &responsePrinter{}
)

type (
	Option func(app *App)

	App struct {
		cli    *cli.App
		client *Client
	}
	Client struct {
		*Config
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
		MaxPushAmount             int64
		MinPushAmount             int64
		MaxDisburseAmount         int64
		MinDisburseAmount         int64
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

func (client *Client) ValidateConfig(requestType RequestType) error {
	pushConfig := client.push.Config
	disburseConfig := client.disburse.Config

	switch requestType {
	case PushPay:
		if pushConfig.PushPayURL == "" ||
			pushConfig.Username == "" ||
			pushConfig.Password == "" ||
			pushConfig.ApiBaseURL == "" ||
			pushConfig.BillerCode == "" ||
			pushConfig.BillerMSISDN == "" ||
			pushConfig.GetTokenURL == "" {
			return ErrNilEmptyConfigVar
		}
		return nil

	case Disburse:

		if disburseConfig.PIN == "" ||
			disburseConfig.BrandID == "" ||
			disburseConfig.RequestURL == "" ||
			disburseConfig.AccountName == "" ||
			disburseConfig.AccountMSISDN == "" {
			return ErrNilEmptyConfigVar
		}

		return nil

	default:
		return nil
	}

}

func Make(config *Config, opts ...Option) *App {

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
		Config:   config,
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

func app(client *Client) *cli.App {

	author := &cli.Author{
		Name:  "Pius Alfred",
		Email: "pmasengwa@techcraft.co.tz",
	}

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "verbose",
			Value: false,
			Usage: "allow verbose output",
		},
	}

	var commands []*cli.Command

	commands = append(
		commands,
		client.MakeLoginCommand(),
		client.MakeConfigCommand(),
		client.MakePushCommand(),
		client.MakeDisburseCommand(),
	)

	a := &cli.App{
		Name:                 "tg",
		Usage:                "command line tool to execute tigopesa push pay and disbursement requests",
		UsageText:            "tg [config | push | disburse ]",
		Version:              "1.0.0",
		Description:          "use this tool to perform push pay requests or disbursement requests.\nall these requests are through tigopesa.\nmake sure the too is correctly configured via config command",
		Commands:             commands,
		Flags:                flags,
		EnableBashCompletion: true,
		Authors: []*cli.Author{
			author,
		},
		Copyright: "Pius Alfred (c) 2021. MIT License",
	}

	return a
}

func (app *App) Run(args []string) error {
	return app.cli.Run(args)
}
