package main

import (
	"log"

	"github.com/evenfound/even-go/node/cmd/evec/app"
	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	app.Init()
	defer app.Close()
	err := app.Run()
	if err != nil {
		if config.Debug {
			serr, ok := err.(stackTracer)
			if ok {
				log.Fatalf("Error: %+v", serr)
			}
		}
		log.Fatal("Error: ", err)
	}
}
