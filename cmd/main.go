package main

import (
	"github.com/techcraftt/tg"
	"log"
	"os"
)

const (
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

var (
	disburseAccountName   string
	disburseAccountMSISDN string
	disburseBrandID       string
	disbursePin           string
	disburseURL           string
	pushUsername          string
	pushBillerMSISDN      string
	pushBIllerCode        string
	pushGetTokenURL       string
	pushURL               string
	pushCallbackURL       string
)

func LoadConfingFromEnv() (*tg.Config, error) {
	disburseAccountName = os.Getenv(envDisburseAccountName)

	return &tg.Config{
		ReferenceIDPrefix:         "",
		AppName:                   "",
		AppVersion:                "",
		ReleaseDate:               "",
		DisburseAccountName:       disburseAccountName,
		DisburseAccountMSISDN:     disburseAccountMSISDN,
		DisburseBrandID:           disburseBrandID,
		DisbursePIN:               disbursePin,
		DisburseRequestURL:        disburseURL,
		PushUsername:              pushUsername,
		PushPassword:              pushUsername,
		PushPasswordGrantType:     "",
		PushApiBaseURL:            "",
		PushGetTokenURL:           "",
		PushBillerMSISDN:          "",
		PushBillerCode:            "",
		PushPayURL:                "",
		PushReverseTransactionURL: "",
		PushHealthCheckURL:        "",
	},nil
}

func main() {

	conf, err := LoadConfingFromEnv()
	if err != nil{
		log.Fatal(err)
	}
	app := tg.Make(conf)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
