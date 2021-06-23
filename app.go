package tg

import (
	"context"
	"github.com/urfave/cli/v2"
)

type (
	Application struct {
		TigoClient interface{}
		VerboseMode bool
		Context context.Context
		CLI *cli.App
	}
)

