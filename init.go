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
	"bufio"
	"fmt"
	"github.com/techcraftt/tg/secret"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func (client *Client) MakeInitCommand() *cli.Command {
	command := &cli.Command{
		Name:         "init",
		Usage:        "initialize the tool",
		Before:       BeforeInitAction,
		After:        client.AfterInitAction,
		Action:       client.OnInitAction,
		OnUsageError: OnInitError,
	}

	return command
}

func OnInitError(context *cli.Context, err error, subcommand bool) error {
	fmt.Printf("error occurred: %v\n", err)
	return nil
}

func (client *Client) OnInitAction(context *cli.Context) error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	if strings.TrimSpace(username) == "" {
		return fmt.Errorf("username is empty, enter a valid one")
	}
	password, err := ReadPassword(10)
	if err != nil {
		return err
	}

	login := &Login{
		Username: strings.TrimSpace(username),
		Password: password,
	}
	client.login = login
	fmt.Print("\nOK\n")
	secret.WithUser(username)(client.secrets)
	_ = client.secrets.Save(password)
	return nil
}

func (client *Client) AfterInitAction(context *cli.Context) error {
	fmt.Printf("username: %s and password: %s\n", client.login.Username, client.login.Password)
	kg, err := client.secrets.Get()
	if err != nil {
		return err
	}
	fmt.Printf("from keyring the details are: %s\n",kg)
	return nil
}

func BeforeInitAction(context *cli.Context) error {
	return nil
}
