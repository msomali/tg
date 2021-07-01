package main

import (
	"github.com/techcraftt/tg"
	"log"
	"os"
)

var (
	envDisburseAccountName   = "TIGO_DISBURSE_ACCOUNT_NAME"
	envDisburseAccountMSISDN = "TIGO_DISBURSE_ACCOUNT_MSISDN"
	envDisburseBrandID       = "TIGO_DISBURSE_BRAND_ID"
	envDisbursePin           = "TIGO_DISBURSE_PIN"
	envDisburseURL           = "TIGO_DISBURSE_URL"
	envPushUsername          = "TIGO_PUSH_USERNAME"
	envPushPassword          = "TIGO_PUSH_PASSWORD"
	envPushBillerMSISDN      = "TIGO_PUSH_BILLER_MSISDN"
	envPushBIllerCode        = "TIGO_PUSH_BILLER_CODE"
	envPushGetTokenURL       = "TIGO_PUSH_TOKEN_URL"
	envPushURL               = "TIGO_PUSH_URL"
	envPushCallbackURL       = "TIGO_PUSH_CALLBACK_URL"
)

func main() {
	app := tg.Make(&tg.Config{})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
