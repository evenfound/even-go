// Copyright (C) 2017-2019 The Even Network Developers

package mbnd

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"runtime"
	"runtime/debug"
)

// winServiceMain is only invoked on Windows.
// It detects when mbnd is running as a service and reacts accordingly.
var winServiceMain func() (bool, error)

// Multi-Blockchain Network Demon is the entry point for start external micro-services.
// Required by defers created in the top-level scope of a main method aren't executed if os.Exit() is called.
func mbnd() error {

	// Load configuration and parse command line.
	// This function also initializes logging and configures it accordingly.
	_, _, err := loadConfig()
	if err != nil {
		return err
	}

	// Get a channel that will be closed when a shutdown signal has been
	// triggered either from an OS signal such as SIGINT (Ctrl+C) or from
	// another subsystem such as the RPC server.
	interrupt := interruptListener()

	// Return now if an interrupt signal was triggered.
	if interruptRequested(interrupt) {
		return nil
	}

	procAttr := &os.ProcAttr{
		Dir:   os.TempDir(),
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   &syscall.SysProcAttr{},
	}

	//@todo in future will move this options to configuration file
	procArgv := []string{
		"--testnet",
	}

	//@todo in future will move this list to configuration file
	bcNets := []string{
		"btcd",
		"ltcd",
	}

	for _, bcNet := range bcNets {

		binFilePath, lookErr := exec.LookPath(bcNet)
		if lookErr != nil {
			panic(lookErr)
		}

		fmt.Printf("Start %s form %v ...\n", bcNet, binFilePath)

		process, err := os.StartProcess(binFilePath, procArgv, procAttr)
		if err != nil {
			panic(err)
		}

		err = process.Release()
		if err != nil {
			panic(err)
		}

		defer process.Kill()

	}

	return nil
}

func Start() {

	// Use all processor cores.
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Block and transaction processing can cause bursty allocations.
	// This limits the garbage collector from excessively overallocating during bursts.
	// This value was arrived at with the help of profiling live usage.
	debug.SetGCPercent(5)

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
	if err := mbnd(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
