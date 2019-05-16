// Copyright (C) 2017-2019 The Even Network Developers

package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/evenfound/even-go/core"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
	"github.com/jessevdk/go-flags"
)

// winServiceMain is only invoked on Windows.
// It detects when mbnd is running as a service and reacts accordingly.
var winServiceMain func() (bool, error)

type (
	// Stop represents stop.
	Stop struct{}
	// Restart represents restart.
	Restart struct{}
	// Options contains global node options.
	Options struct {
		Version bool `short:"v" long:"version" description:"Print the version number and exit"`
	}
)

var (
	stopServer    Stop
	restartServer Restart
	options       Options
	parser        = flags.NewParser(&options, flags.HelpFlag|flags.PassDoubleDash)
)

func main() {

	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Block and transaction processing can cause bursty allocations.
	// This limits the garbage collector from excessively overallocating during bursts.
	// This value was arrived at with the help of profiling live usage.
	debug.SetGCPercent(10)

	// Call serviceMain on Windows to handle running as a service.
	// When the return isService flag is true, exit now since we ran as a service.
	// Otherwise, just fall through to normal operation.
	if runtime.GOOS == "windows" {
		isService, err := winServiceMain()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isService {
			os.Exit(0)
		}
	}

	// Work around defer not working after os.Exit()
	if err := worker(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func worker() error {

	// Load configuration and parse command line.
	// This function also initializes logging and configures it accordingly.
	_, _, err := loadConfig()
	if err != nil {
		return err
	}

	for name, spec := range commandList {
		if _, err := parser.AddCommand(name, spec.shortDescription, spec.longDescription, spec.command); err != nil {
			return err
		}
	}

	if _, err := parser.Parse(); err != nil {
		return err
	}

	return nil
}
