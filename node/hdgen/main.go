package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	wg     = sync.WaitGroup{}
	logger *log.Logger
)

func main() {
	// Recovering all errors during the process
	defer errorHandler()

	wg.Add(2)

	go RPCConnect()

	fmt.Println("Listening for RPC   127.0.0.1:" + config.rpcPort)

	go HTTPConnect()

	fmt.Println("Listening for HTTP  127.0.0.1:" + config.httpPort)

	wg.Wait()
}

// Catch errors during process
func errorHandler() {
	if rec := recover(); rec != nil {
		err := rec.(error)
		logger.Printf("Unhandled error: %v\n", err.Error())
		fmt.Fprintf(os.Stderr, "Program quit unexpectedly; please check your logs\n")
		os.Exit(1)
	}
}
