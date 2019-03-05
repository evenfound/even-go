package main

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/user"
)

type ServiceConfig struct {
	rpcPort  string
	httpPort string
	dataDir  string
}

var config = ServiceConfig{}

// Getting current user home directory
func getDataDir() string {
	var user, err = user.Current()

	if err != nil {
		panic(errors.New(" Can not recognize current user "))
	}

	return user.HomeDir
}

func init() {

	var httpPort = flag.String("httpport", "9080", "Listening for HTTP connections on port <port>")
	var rpcPort = flag.String("rpcport", "9081", "Listening for RPC connections on port <port>")
	var dataDir = flag.String("datadir", getDataDir(), "Specify data directory")

	flag.Parse()

	if rpcPort == httpPort {
		panic(errors.New("CONFLICT : Can not listen http/rpc server"))
	}

	config.rpcPort = *rpcPort
	config.httpPort = *httpPort
	config.dataDir = *dataDir

	f, err := os.Create("app/log")

	if err != nil {
		panic(err)
	}

	logger = log.New(f, "", log.LstdFlags)
}
