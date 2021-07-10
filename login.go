package tg

import (
	"context"
	"fmt"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
	"time"
)

func (client *Client) MakeLoginCommand() *cli.Command {
	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "auto",
			Usage: "allow automatic login",
			Value: false,
		},
	}
	command := &cli.Command{
		Name:         "login",
		Usage:        "login to use the tool",
		Before:       client.BeforeLoginAction,
		After:        client.AfterLoginAction,
		Action:       client.OnLoginAction,
		OnUsageError: client.OnLoginError,
		Flags:        flags,
	}

	return command
}

func (client *Client) BeforeLoginAction(ctx *cli.Context) error {

	//todo: validate number (msisdn)
	//todo: validate amount
	//todo: check all parameters
	//todo: prompt confirmation message
	return nil
}

func (client *Client) OnLoginAction(ctx *cli.Context) error {
	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	msisdn := ctx.String("phone")
	id := ctx.String("id")
	remarks := ctx.String("remarks")
	amount := ctx.Float64("amount")

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

func (client *Client) AfterLoginAction(ctx *cli.Context) error {
	return nil
}

func (client *Client) OnLoginError(context *cli.Context, err error, isSubcommand bool) error {
	return nil
}
