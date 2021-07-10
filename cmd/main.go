package main

import (
	"github.com/techcraftt/tg"
	"log"
	"os"
)


func main() {
	conf, err := tg.LoadConfFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	app := tg.Make(conf)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
