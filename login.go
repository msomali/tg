/*
 * MIT License
 *
 * Copyright (c) 2021 TECHCRAFT TECHNOLOGIES CO LTD.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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
