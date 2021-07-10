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
	"fmt"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

const (
	envDisburseAccountName   = "TIGO_DISBURSE_ACCOUNT_NAME"
	envDisburseAccountMSISDN = "TIGO_DISBURSE_ACCOUNT_MSISDN"
	envDisburseBrandID       = "TIGO_DISBURSE_BRAND_ID"
	envDisbursePin           = "TIGO_DISBURSE_PIN"
	envDisburseURL           = "TIGO_DISBURSE_URL"
	envPushUsername          = "TIGO_PUSH_USERNAME"
	envPushPassword          = "TIGO_PUSH_PASSWORD"
	envPushBaseURL           = "TIGO_PUSH_BASE_URL"
	envPushBillerMSISDN      = "TIGO_PUSH_BILLER_MSISDN"
	envPushBillerCode        = "TIGO_PUSH_BILLER_CODE"
	envPushGetTokenURL       = "TIGO_PUSH_TOKEN_URL"
	envPushPayURL            = "TIGO_PUSH_PAY_URL"
	envPushMaxAmount         = "TIGO_PUSH_MAX_AMOUNT"
	envPushMinAmount         = "TIGO_PUSH_MIN_AMOUNT"
	envDisburseMaxAmount     = "TIGO_DISBURSE_MAX_AMOUNT"
	envDisburseMinAmount     = "TIGO_DISBURSE_MIN_AMOUNT"
)

func LoadConfFromEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file %v\n", err)
	}

	pushMaxAmount := os.Getenv(envPushMaxAmount)
	pushMinAmount := os.Getenv(envPushMinAmount)
	disburseMaxAmount := os.Getenv(envDisburseMaxAmount)
	disburseMinAmount := os.Getenv(envDisburseMinAmount)

	disburseName := os.Getenv(envDisburseAccountName)
	disburseMSISDN := os.Getenv(envDisburseAccountMSISDN)
	disburseBrandID := os.Getenv(envDisburseBrandID)
	disbursePIN := os.Getenv(envDisbursePin)
	disburseURL := os.Getenv(envDisburseURL)

	pushName := os.Getenv(envPushUsername)
	pushPassword := os.Getenv(envPushPassword)
	pushBaseURL := os.Getenv(envPushBaseURL)
	pushTokenURL := os.Getenv(envPushGetTokenURL)
	pushPayURL := os.Getenv(envPushPayURL)
	pushMSISDN := os.Getenv(envPushBillerMSISDN)
	pushBillerCode := os.Getenv(envPushBillerCode)

	intPushMaxAmount, err := strconv.ParseInt(pushMaxAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	intPushMinAmount, err := strconv.ParseInt(pushMinAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	intDisburseMaxAmount, err := strconv.ParseInt(disburseMaxAmount, 10, 64)
	if err != nil {
		return nil, err
	}

	intDisburseMinAmount, err := strconv.ParseInt(disburseMinAmount, 10, 64)
	if err != nil {
		return nil, err
	}
	return &Config{
		ReferenceIDPrefix:         "TGCLI",
		AppName:                   "TG",
		AppVersion:                "v1.0.0",
		ReleaseDate:               "JULY 2021",
		MaxPushAmount:             intPushMaxAmount,
		MinPushAmount:             intPushMinAmount,
		MaxDisburseAmount:         intDisburseMaxAmount,
		MinDisburseAmount:         intDisburseMinAmount,
		DisburseAccountName:       disburseName,
		DisburseAccountMSISDN:     disburseMSISDN,
		DisburseBrandID:           disburseBrandID,
		DisbursePIN:               disbursePIN,
		DisburseRequestURL:        disburseURL,
		PushUsername:              pushName,
		PushPassword:              pushPassword,
		PushPasswordGrantType:     "password",
		PushApiBaseURL:            pushBaseURL,
		PushGetTokenURL:           pushTokenURL,
		PushBillerMSISDN:          pushMSISDN,
		PushBillerCode:            pushBillerCode,
		PushPayURL:                pushPayURL,
		PushReverseTransactionURL: "nil",
		PushHealthCheckURL:        "nil",
	}, nil
}

func (client *Client) MakeConfigCommand() *cli.Command {

	config := &cli.Command{
		Name:        "config",
		Usage:       "configure the tool",
		UsageText:   "tg config --push --file=file.yaml ",
		Description: "run this command to configure the tools to use credentials provided by tigo pesa in integration stage",
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
		Subcommands: nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "file",
				Value: "config.yaml",
				Usage: "specify file with config",
			},

			&cli.BoolFlag{
				Name:  "push",
				Usage: "configuration for push pay",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "disburse",
				Usage: "configuration for disburse",
				Value: false,
			},
		},
	}

	return config
}
