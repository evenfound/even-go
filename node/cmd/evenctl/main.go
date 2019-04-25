// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/evenfound/even-go/node/cmd/evenctl/app"
	"github.com/evenfound/even-go/node/cmd/evenctl/config"
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
