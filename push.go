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
			Name:    "phone",
			Usage:   "phone number to receive request",
			Value:   "",
		},
		&cli.StringFlag{
			Name:  "remarks",
			Value: "",
			Usage: "the description of the request",
		},
		&cli.StringFlag{
			Name:    "id",
			Value:   "",
			Usage:   "the reference id of the request",
		},
		&cli.Float64Flag{
			Name:  "amount",
			Value: 0.0,
			Usage: "amount to be paid, from request",
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
	id := ctx.String("id")
	remarks := ctx.String("remarks")
	amount := ctx.Float64("amount")

	request := push.PayRequest{
		CustomerMSISDN: msisdn,
		BillerMSISDN:   client.push.BillerMSISDN,
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
	return nil
}
