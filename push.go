package tg

import (
	"context"
	"fmt"
	"github.com/techcraftt/tigosdk/push"
	"github.com/urfave/cli/v2"
	"time"
)

func (app *App) MakePushCommand() *cli.Command {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:    "phone",
			Aliases: []string{"msisdn"},
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
			Aliases: []string{"reference"},
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
		Aliases:      []string{"p"},
		Usage:        "send push pay request",
		Before:       app.BeforePushAction,
		After:        app.AfterPushAction,
		Action:       app.OnPushAction,
		OnUsageError: app.OnPushError,
		Flags:        flags,
	}

	return command
}

func (app *App) BeforePushAction(ctx *cli.Context) error {

	//todo: validate number (msisdn)
	//todo: validate amount
	//todo: check all parameters
	//todo: prompt confirmation message
	return nil
}

func (app *App) OnPushAction(ctx *cli.Context) error {
	ctx2, cancel := context.WithTimeout(ctx.Context, time.Minute)
	defer cancel()

	msisdn := ctx.String("phone")
	id := ctx.String("id")
	remarks := ctx.String("remarks")
	amount := ctx.Float64("amount")

	request := push.PayRequest{
		CustomerMSISDN: msisdn,
		BillerMSISDN:   app.push.BillerMSISDN,
		Amount:         int(amount),
		Remarks:        remarks,
		ReferenceID:    id,
	}
	response, err := app.push.Pay(ctx2, request)
	fmt.Printf("response: %v\n", response)
	return err
}

func (app *App) AfterPushAction(ctx *cli.Context) error {
	return nil
}

func (app *App) OnPushError(context *cli.Context, err error, isSubcommand bool) error {
	return nil
}
