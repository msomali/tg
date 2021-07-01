package tg

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

func (client *Client) MakeDisburseCommand() *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "phone",
			Usage:   "phone number to receive request",
			Value:   "",
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
		Name:         "disburse",
		Usage:        "send disburse pay request",
		Before:       client.BeforeDisburseAction,
		After:        client.AfterDisburseAction,
		Action:       client.OnDisburseAction,
		OnUsageError: client.OnDisburseError,
		Flags:        flags,
	}

	return command
}

func (client *Client) BeforeDisburseAction(ctx *cli.Context) error {

	//todo: validate number (msisdn)
	//todo: validate amount
	//todo: check all parameters
	//todo: prompt confirmation message
	return nil
}

func (client *Client) OnDisburseAction(ctx *cli.Context) error {
	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	msisdn := ctx.String("phone")
	id := ctx.String("id")
	amount := ctx.Float64("amount")

	response, err := client.disburse.Disburse(ctx2, id,msisdn,amount)
	fmt.Printf("response: %v\n", response)
	return err
}

func (client *Client) AfterDisburseAction(ctx *cli.Context) error {
	return nil
}

func (client *Client) OnDisburseError(context *cli.Context, err error, isSubcommand bool) error {
	return nil
}
