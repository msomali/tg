package tg

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)



func main() {
	app := &cli.App{
		Name:  "tg",
		Usage: "command line tool for pushpay and disbursement via tigopesa",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func App() *cli.App {
	return &cli.App{}
}
