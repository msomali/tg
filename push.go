package tg

import (
	"context"
	"errors"
	"fmt"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
	"regexp"
	"strings"
	"time"
)

var (
	errInvalidPhoneNumber = errors.New("invalid phone number format: allowed format: 255XXXXXXXXX OR 0XXXXXXXXX")
)

func (client *Client) MakePushCommand() *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "phone",
			Usage: "phone number to receive request",
			Value: "",
			Required: true,
		},
		&cli.StringFlag{
			Name:  "remarks",
			Value: "",
			Usage: "the description of the request",
			Required: true,
		},
		&cli.Float64Flag{
			Name:  "amount",
			Value: 0.0,
			Usage: "amount to be paid, from request",
			Required: true,
		},
	}
	command := &cli.Command{
		Name:         "push",
		Usage:        "send push pay request",
		Before:       client.BeforePushAction,
		After:        client.AfterPushAction,
		Action:       client.OnPushAction,
		OnUsageError: client.OnPushError,
		Flags:        flags,
	}

	return command
}

func (client *Client) BeforePushAction(ctx *cli.Context) error {
	remarks := ctx.String("remarks")
	phone := ctx.String("phone")
	amount := int64(ctx.Float64("amount"))
	verbose := ctx.Bool("verbose")
	defer func(v bool) {
		if v{
			fmt.Printf("perfoming push pay request described as \"%s\" of amount: %d to phone: %s\n",
				remarks, amount, phone)
		}
	}(verbose)

	err := CheckPhoneNumber(phone)
	if err != nil{
		return err
	}
	if amount< client.MinPushAmount || amount > client.MaxPushAmount{
		return fmt.Errorf("the amount (%d) is out of range: allowed MAX is %d, allowed MIN is %d\n", amount,client.MaxPushAmount, client.MinPushAmount)
	}
	err = client.ValidateConfig(PushPay)
	if err != nil {
		return fmt.Errorf("check your push configs: %s", err.Error())
	}
	return nil
}

func (client *Client) OnPushAction(ctx *cli.Context) error {
	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()
	msisdn := ctx.String("phone")
	remarks := ctx.String("remarks")
	amount := ctx.Float64("amount")
	id := fmt.Sprintf("%s%s%d", client.push.Config.BillerCode,client.ReferenceIDPrefix,time.Now().Local().Unix())

	request := push.PayRequest{
		CustomerMSISDN: msisdn,
		Amount:         int(amount),
		Remarks:        remarks,
		ReferenceID:    id,
	}
	response, err := client.push.Pay(ctx2, request)
	fmt.Printf("response: %v\n", response)
	return err
}

func (client *Client) AfterPushAction(ctx *cli.Context) error {
	return nil
}

func (client *Client) OnPushError(context *cli.Context, err error, isSubcommand bool) error {
	fmt.Printf("error while push pay %v\n", err)
	return nil
}

func CheckPhoneNumber(phone string)error{
	strLen := len(phone)
	//check length
	if strLen != 10 && strLen != 12{
		return errInvalidPhoneNumber
	}

	re := regexp.MustCompile("^\\d+$")

	match := re.MatchString(phone)

	if ! match{
		return fmt.Errorf("%v: letters are not allowed\n",errInvalidPhoneNumber)
	} else{
		if strLen == 12{
			//if the length == 12
			//it should start with 255
			if !strings.HasPrefix(phone, "255"){
				return fmt.Errorf("%v: should start with \"255\"\n",errInvalidPhoneNumber)
			}

		}
		if strLen == 10{
			//if len is 10
			//it should start with 0
			if !strings.HasPrefix(phone, "0"){
				return fmt.Errorf("%v: should start with \"0\"\n",errInvalidPhoneNumber)
			}
		}
		return nil
	}
}
