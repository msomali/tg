package main

import (
	"context"
	"github.com/urfave/cli/v2"
)

func InitCommand(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:         "init",
		Usage:        "initialize",
		UsageText:    "initialize the tool including setting up the tigopesa integration details",
		Description:  "This command is used to configure ne",
		Before:       beforeInitFunc,
		After:        afterInitFunc,
		Action:       executeInit,
		OnUsageError: onInitError,
	}
}

// beforeInitFunc checks if the app was initialized before this command so as to warn the
// user on any potential of deleting the previous configs
func beforeInitFunc(ctx *cli.Context) error {
	return nil
}

// afterInitFunc checks if the initialized configurations has been stored
//correctly for later usage including printing them to user for verification
func afterInitFunc(ctx *cli.Context) error {
	return nil
}

func executeInit(ctx *cli.Context) error {
	return nil
}

func onInitError(context *cli.Context, err error, isSubcommand bool) error {
	return nil
}
