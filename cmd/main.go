package main

import (
	"github.com/techcraftt/tg"
	"log"
	"os"
)



func main() {
	app := tg.Make(&tg.Config{})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
