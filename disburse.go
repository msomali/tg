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
			Name:  "phone",
			Usage: "phone number to receive request",
			Value: "",
		},

		&cli.StringFlag{
			Name:  "id",
			Value: "",
			Usage: "the reference id of the request",
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

	phone := ctx.String("phone")
	amount := int64(ctx.Float64("amount"))
	verbose := ctx.Bool("verbose")
	defer func(v bool) {
		if v {
			fmt.Printf("perfoming disbursement of amount: %d to phone: %s\n",
				amount, phone)
		}
	}(verbose)

	err := CheckPhoneNumber(phone)
	if err != nil {
		return err
	}
	if amount < client.MinDisburseAmount || amount > client.MaxDisburseAmount {
		return fmt.Errorf("the amount (%d) is out of range: allowed MAX is %d, allowed MIN is %d\n", amount, client.MaxPushAmount, client.MinPushAmount)
	}
	err = client.ValidateConfig(Disburse)
	if err != nil {
		return fmt.Errorf("check your disburse configs: %s", err.Error())
	}
	return nil
}

func (client *Client) OnDisburseAction(ctx *cli.Context) error {

	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	msisdn := ctx.String("phone")
	id := fmt.Sprintf("%s%s%d", client.disburse.BrandID, client.ReferenceIDPrefix, time.Now().Local().Unix())
	amount := ctx.Float64("amount")
	response, err := client.disburse.Disburse(ctx2, id, msisdn, amount)
	defaultResponsePrinter.TextOut(Disburse, response)
	return err
}

func (client *Client) AfterDisburseAction(ctx *cli.Context) error {
	return nil
}

func (client *Client) OnDisburseError(context *cli.Context, err error, isSubcommand bool) error {
	fmt.Printf("error during disbursement %v\n", err)
	return nil
}
