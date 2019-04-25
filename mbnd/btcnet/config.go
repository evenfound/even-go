// Copyright (C) 2017-2019 The Even Network Developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package btcnet

import (
	"fmt"
	network "net"
	"os"
	"path/filepath"
	"runtime"

	"github.com/btcsuite/btclog"
	"github.com/btcsuite/btcutil"
	_ "github.com/evenfound/even-go/mbnd/common/ffldb"
	"github.com/evenfound/even-go/utility"
	"github.com/jessevdk/go-flags"
)

type (
	// serviceOptions defines the configuration options for the daemon as a service on Windows.
	serviceOptions struct {
		ServiceCommand string `short:"s" long:"service" description:"Service command {install, remove, start, stop}"`
	}

	// config defines the configuration options for mbnd.
	// See loadConfig for details on the configuration load process.
	config struct {
		ShowVersion       bool     `short:"V" long:"version" description:"Display version information and exit"`
		ConfigFile        string   `short:"C" long:"configfile" description:"Path to configuration file"`
		LogDir            string   `long:"logdir" description:"Directory to logger output."`
		DataDir           string   `short:"b" long:"datadir" description:"Directory to store data"`
		ExternalDir       string   `short:"e" long:"externaldir" description:"Directory to store data"`
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
		DebugLevel        string   `short:"d" long:"debuglevel" description:"Logging level for all subsystems {trace, debug, info, warn, error, critical} -- You may also specify <subsystem>=<level>,<subsystem2>=<level>,... to set the logger level for individual subsystems -- Use show to list available subsystems"`
		MiningAddrs       []string `long:"miningaddr" description:"Add the specified payment address to the list of addresses to use for generated blocks -- At least one address is required if the generate option is set"`
		UserAgentComments []string `long:"uacomment" description:"Comment to add to the user agent -- See BIP 14 for more information."`
		BlocksOnly        bool     `long:"blocksonly" description:"Do not accept transactions from remote peers."`
		TxIndex           bool     `long:"txindex" description:"Maintain a full hash-based transaction index which makes all transactions available via the getrawtransaction RPC"`
		DropTxIndex       bool     `long:"droptxindex" description:"Deletes the hash-based transaction index from the database on start up and then exits."`
		AddrIndex         bool     `long:"addrindex" description:"Maintain a full address-based transaction index which makes the searchrawtransactions RPC available"`
		DropAddrIndex     bool     `long:"dropaddrindex" description:"Deletes the address-based transaction index from the database on start up and then exits."`
		lookup            func(string) ([]network.IP, error)
		miningAddrs       []btcutil.Address
		minRelayTxFee     btcutil.Amount
		whitelists        []*network.IPNet
	}
)

const (
	defaultConfigFilename  = "mbnd.conf"
	defaultLogLevel        = "info"
	defaultLogDirname      = "logs"
	defaultDataDirname     = "data"
	defaultExternalDirname = "external"
	defaultDbType          = "ffldb"
	blockDbNamePrefix      = "blocks"
)

var (
	// Init and setup logger.
	logger = btclog.NewBackend(os.Stdout).Logger("MBND")

	// runServiceCommand is only set to a real function on Windows.
	// It is used to parse and execute service commands specified via the -s flag.
	runServiceCommand func(string) error

	defaultHomeDir     = utility.AppDataDir(".evenet", false)
	defaultLogDir      = filepath.Join(defaultHomeDir, defaultLogDirname)
	defaultDataDir     = filepath.Join(defaultHomeDir, defaultDataDirname)
	defaultExternalDir = filepath.Join(defaultHomeDir, defaultDataDirname, defaultExternalDirname)
	defaultConfigFile  = filepath.Join(defaultHomeDir, defaultConfigFilename)
	defaultRPCKeyFile  = filepath.Join(defaultHomeDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(defaultHomeDir, "rpc.cert")
)

// newConfigParser returns a new command line flags parser.
func newConfigParser(cfg *config, so *serviceOptions, options flags.Options) *flags.Parser {
	parser := flags.NewParser(cfg, options)
	if runtime.GOOS == "windows" {
		_, _ = parser.AddGroup("Service Options", "Service Options", so)
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
// The above results in mbnd functioning properly without any config settings while still allowing the user
// to override settings with config files and command line options.
// Command line options always take precedence.
func loadConfig() error {

	// Check if configuration is already loaded
	if cfg != nil {
		return nil
	}

	// Default config.
	preCfg := config{
		ConfigFile:  defaultConfigFile,
		DebugLevel:  defaultLogLevel,
		LogDir:      defaultLogDir,
		DataDir:     defaultDataDir,
		ExternalDir: defaultExternalDir,
		DbType:      defaultDbType,
		RPCKey:      defaultRPCKeyFile,
		RPCCert:     defaultRPCCertFile,
		TxIndex:     true,
		AddrIndex:   true,
	}

	// Service options which are only added on Windows.
	serviceOpts := serviceOptions{}

	// Pre-parse the command line options to see if an alternative config file or the version flag was specified.
	// Any errors aside from the help message error can be ignored here
	// since they will be caught by the final parse below.
	preParser := newConfigParser(&preCfg, &serviceOpts, flags.HelpFlag)

	_, err := preParser.Parse()
	if err != nil {
		if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return err
		}
	}

	// Perform service command and exit if specified.
	// Invalid service commands show an appropriate error.
	// Only runs on Windows since the runServiceCommand function will be nil when not on Windows.
	if serviceOpts.ServiceCommand != "" && runServiceCommand != nil {
		err := runServiceCommand(serviceOpts.ServiceCommand)
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(0)
	}

	cfg = &preCfg

	return nil
}
