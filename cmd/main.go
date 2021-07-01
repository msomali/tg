package main

import (
	"github.com/techcraftt/tg"
	"log"
	"os"
)

var (
	disburseAccountName   = "TIGO_DISBURSE_ACCOUNT_NAME"
	disburseAccountMSISDN = "TIGO_DISBURSE_ACCOUNT_MSISDN"
	disburseBrandID       = "TIGO_DISBURSE_BRAND_ID"
	disbursePin           = "TIGO_DISBURSE_PIN"
	disburseURL           = "TIGO_DISBURSE_URL"
	pushUsername          = "TIGO_PUSH_USERNAME"
	pushPassword          = "TIGO_PUSH_PASSWORD"
	pushBillerMSISDN      = "TIGO_PUSH_BILLER_MSISDN"
	pushBIllerCode        = "TIGO_PUSH_BILLER_CODE"
	pushGetTokenURL       = "TIGO_PUSH_TOKEN_URL"
	pushURL               = "TIGO_PUSH_URL"
	pushCallbackURL       = "TIGO_PUSH_CALLBACK_URL"
)

func main() {
	app := tg.Make(&tg.Config{})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
