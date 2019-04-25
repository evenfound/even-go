// Copyright (C) 2017-2019 The Even Network Developers

package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"

	"github.com/btcsuite/btclog"
	"github.com/evenfound/even-go/node/utility"
	"github.com/jessevdk/go-flags"
)

type (
	// config defines the configuration options for evenet.
	// See loadConfig for details on the configuration load process.
	config struct {
		ShowVersion       bool     `short:"V" long:"version" description:"Display version information and exit"`
		ConfigFile        string   `short:"C" long:"configfile" description:"Path to configuration file"`
		DataDir           string   `short:"b" long:"datadir" description:"Directory to store data"`
		LogDir            string   `long:"logdir" description:"Directory to log output."`
		RPCUser           string   `short:"u" long:"rpcuser" description:"Username for RPC connections"`
		RPCPass           string   `short:"P" long:"rpcpass" default-mask:"-" description:"Password for RPC connections"`
		RPCLimitUser      string   `long:"rpclimituser" description:"Username for limited RPC connections"`
		RPCLimitPass      string   `long:"rpclimitpass" default-mask:"-" description:"Password for limited RPC connections"`
		RPCCert           string   `long:"rpccert" description:"File containing the certificate file"`
		RPCKey            string   `long:"rpckey" description:"File containing the certificate key"`
		Proxy             string   `long:"proxy" description:"Connect via SOCKS5 proxy (eg. 127.0.0.1:9050)"`
		ProxyUser         string   `long:"proxyuser" description:"Username for proxy server"`
		ProxyPass         string   `long:"proxypass" default-mask:"-" description:"Password for proxy server"`
		TestNet3          bool     `long:"testnet" description:"Use the test network"`
		RegressionTest    bool     `long:"regtest" description:"Use the regression test network"`
		SimNet            bool     `long:"simnet" description:"Use the simulation test network"`
		DbType            string   `long:"dbtype" description:"Database backend to use for the Block Chain"`
		CPUProfile        string   `long:"cpuprofile" description:"Write CPU profile to the specified file"`
		DebugLevel        string   `short:"d" long:"debuglevel" description:"Logging level for all subsystems {trace, debug, info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the log level for individual subsystems -- Use show to list available subsystems"`
		MiningAddrs       []string `long:"miningaddr" description:"Add the specified payment address to the list of addresses to use for generated blocks -- At least one address is required if the generate option is set"`
		UserAgentComments []string `long:"uacomment" description:"Comment to add to the user agent -- See BIP 14 for more information."`
		BlocksOnly        bool     `long:"blocksonly" description:"Do not accept transactions from remote peers."`
		TxIndex           bool     `long:"txindex" description:"Maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC"`
		DropTxIndex       bool     `long:"droptxindex" description:"Deletes the hash-based transaction index from the database on start up and then exits."`
		AddrIndex         bool     `long:"addrindex" description:"Maintain a full address-based transaction index which makes the searchrawtransactions RPC available"`
		DropAddrIndex     bool     `long:"dropaddrindex" description:"Deletes the address-based transaction index from the database on start up and then exits."`
		lookup            func(string) ([]net.IP, error)
		whitelists        []*net.IPNet
	}

	// serviceOptions defines the configuration options for the daemon as a service on Windows.
	serviceOptions struct {
		ServiceCommand string `short:"s" long:"service" description:"Service command {install, remove, start, stop}"`
	}
)

const (
	defaultConfigFilename = "evenet.conf"
	defaultDataDirname    = "data"
	defaultLogDirname     = "logs"
	defaultLogLevel       = "info"
	defaultDbType         = "ffldb"
)

var (
	defaultHomeDir     = utility.AppDataDir(".evenet", false)
	defaultConfigFile  = filepath.Join(defaultHomeDir, defaultConfigFilename)
	defaultDataDir     = filepath.Join(defaultHomeDir, defaultDataDirname)
	defaultRPCKeyFile  = filepath.Join(defaultHomeDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(defaultHomeDir, "rpc.cert")
	defaultLogDir      = filepath.Join(defaultHomeDir, defaultLogDirname)

	// runServiceCommand is only set to a real function on Windows.
	// It is used to parse and execute service commands specified via the -s flag.
	runServiceCommand func(string) error

	// Init and setup logger.
	logger = btclog.NewBackend(os.Stdout).Logger("EVEN")
)

// newConfigParser returns a new command line flags parser.
func newConfigParser(cfg *config, so *serviceOptions, options flags.Options) *flags.Parser {

	parser := flags.NewParser(cfg, options)

	if runtime.GOOS == "windows" {
		_, err := parser.AddGroup("Service Options", "Service Options", so)
		if err != nil {
			return nil
		}
	}

	return parser
}

// loadConfig initializes and parses the config using a config file and command line options.
//
// The configuration proceeds as follows:
// 	1) Start with a default config with sane settings
// 	2) Pre-parse the command line to check for an alternative config file
// 	3) Load configuration file overwriting defaults with any specified options
// 	4) Parse CLI options and overwrite/add any specified options
//
// The above results in evenet functioning properly without any config settings while still allowing the user to
// override settings with config files and command line options.  Command line options always take precedence.
func loadConfig() (*config, []string, error) {

	// Default config.
	cfg := config{
		ConfigFile: defaultConfigFile,
		DebugLevel: defaultLogLevel,
		DataDir:    defaultDataDir,
		LogDir:     defaultLogDir,
		DbType:     defaultDbType,
		RPCKey:     defaultRPCKeyFile,
		RPCCert:    defaultRPCCertFile,
	}

	// Service options which are only added on Windows.
	serviceOpts := serviceOptions{}

	// Pre-parse the command line options to see if an alternative config file or the version flag was specified.
	// Any errors aside from the help message error can be ignored here since they will be caught by the final parse below.
	preCfg := cfg
	preParser := newConfigParser(&preCfg, &serviceOpts, flags.HelpFlag)

	_, err := preParser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			fmt.Fprintln(os.Stderr, err)
			return nil, nil, err
		}
	}

	// Perform service command and exit if specified.
	// Invalid service commands show an appropriate error.
	// Only runs on Windows since the runServiceCommand function will be nil when not on Windows.
	if serviceOpts.ServiceCommand != "" && runServiceCommand != nil {

		err := runServiceCommand(serviceOpts.ServiceCommand)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		os.Exit(0)
	}

	return &cfg, nil, nil
}
