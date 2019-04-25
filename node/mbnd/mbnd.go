// Copyright (C) 2017-2019 The Even Network Developers

package mbnd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"syscall"
)

// winServiceMain is only invoked on Windows.
// It detects when mbnd is running as a service and reacts accordingly.
var winServiceMain func() (bool, error)

// External start the Multi-Blockchain Network Demon.
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

// Wrapper run the Multi-Blockchain Network Demon.
func run(name string, argv []string, attr *os.ProcAttr) {

	process, err := os.StartProcess(name, argv, attr)
	if err != nil {
		panic(err)
	}

	err = process.Release()
	if err != nil {
		panic(err)
	}

	defer func() {
		_ = process.Kill()
	}()
}

// Multi-Blockchain Network Demon is the entry point for start external micro-services.
// Required by defers created in the top-level scope of a main method aren't executed if os.Exit() is called.
func mbnd() error {

	// Load configuration and parse command line.
	// This function also initializes logging and configures it accordingly.
	err := loadConfig()
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

	//@todo in future will move this options to configuration file
	argv := []string{
		"--testnet",
		"--txindex",
		"--addrindex",
	}

	//@todo in future will move this list to configuration file
	networks := []string{
		"btcd",
		"ltcd",
	}

	procAttr := &os.ProcAttr{
		Dir:   os.TempDir(),
		Env:   os.Environ(),
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Sys:   &syscall.SysProcAttr{},
	}

	for _, net := range networks {

		binFilePath, lookErr := exec.LookPath(net)
		if lookErr != nil {
			panic(lookErr)
		}

		logger.Infof("Start %s form %v ...", net, binFilePath)

		procArgv := append(argv,
			"--logdir="+filepath.Join(cfg.ExternalDir, net, defaultLogDirname),
			"--datadir="+filepath.Join(cfg.ExternalDir, net, defaultDataDirname))

		go run(binFilePath, procArgv, procAttr)
	}

	defer func() {
		_ = os.Stdout.Sync()
	}()

	// Wait until the interrupt signal is received from an OS signal or shutdown
	// is requested through one of the subsystems such as the RPC server.
	<-interrupt

	return nil
}
