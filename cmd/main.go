package main

import (
	"github.com/techcraftt/tg"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)



func main() {
	app := tg.NewApplication()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func App() *cli.App {
	return &cli.App{}
}
