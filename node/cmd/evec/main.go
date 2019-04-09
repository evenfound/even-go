package main

import (
	"log"

	"github.com/evenfound/even-go/node/cmd/evec/app"
	"github.com/evenfound/even-go/node/cmd/evec/config"

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
