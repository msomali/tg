package tg

import (
	"context"
	"fmt"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
	"time"
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
	phone := ctx.String("phone")
	amount := int64(ctx.Float64("amount"))
	if amount< client.MinPushAmount || amount > client.MaxPushAmount{
		return fmt.Errorf("the amount (%d) is out of range: allowed MAX is %d, allowed MIN is %d\n", amount,client.MaxPushAmount, client.MinPushAmount)
	}
	remarks := ctx.String("remarks")
	if ctx.Bool("verbose") {
		fmt.Printf("perfoming push pay request described as \"%s\" of amount: %d to phone: %s\n",
			remarks, amount, phone)
	}
	err := client.ValidateConfig(PushPay)
	if err != nil {
		return fmt.Errorf("check your push configs: %s", err.Error())
	}

	//todo: validate number (msisdn)
	//todo: validate amount
	//todo: check all parameters
	//todo: prompt confirmation message
	return nil
}

func (client *Client) OnPushAction(ctx *cli.Context) error {
	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()
	msisdn := ctx.String("phone")
	remarks := ctx.String("remarks")
	amount := ctx.Float64("amount")
	id := fmt.Sprintf("%s%d", client.push.Config.BillerCode, time.Now().Local().Unix())

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
