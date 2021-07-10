package main

import (
	"fmt"
	"github.com/techcraftt/tg"
	"log"
	"os"
)

func main() {
	config, err := tg.LoadConfFromEnv()
	if err != nil {
		fmt.Printf("error while loading config %v\n", err)
		config = &tg.Config{}
	}
	app := tg.Make(config)
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
