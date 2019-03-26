package main

import (
	"github.com/evenfound/even-go/node/cmd/evenctl/app"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"log"

	"github.com/ztrue/tracerr"
)

func main() {
	app.Init()
	defer app.Close()
	err := app.Run()
	if err != nil {
		if config.Debug {
			tracerr.PrintSourceColor(err)
		}
		log.Fatal("Fatal error: ", err)
	}
}
