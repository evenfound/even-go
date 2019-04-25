// Copyright (C) 2017-2019 The Even Network Developers

package main

import (
	"github.com/evenfound/even-go/node/cmd"
)

type commandItem struct {
	shortDescription string
	longDescription  string
	command          interface{}
}

var commandList = map[string]commandItem{

	"version": {
		"Even Network core version",
		"Display current Even Network core version",
		&cmd.CoreVersion{},
	},

	"gencerts": {
		"Generate certificates",
		"Generate self-signed certificates",
		&cmd.GenerateCertificates{},
	},

	"init": {
		"initialize a new repo and exit",
		"Initializes a new repo without starting the server",
		&cmd.Init{},
	},

	"status": {
		"get the repo status",
		"Returns the status of the repo ― Uninitialized, Encrypted, Decrypted. Also returns whether Tor is available.",
		&cmd.Status{},
	},

	"setapicreds": {
		"set API credentials",
		"The API password field in the config file takes a SHA256 hash of the password. This command will generate the hash for you and save it to the config file.",
		&cmd.SetAPICreds{},
	},

	"start": {
		"start the EvenNetwork-Server",
		"The start command starts the EvenNetwork-Server",
		&cmd.Start{},
	},

	"stop": {
		"shutdown the server and disconnect",
		"The stop command disconnects from peers and shuts down EvenNetwork-Server",
		&stopServer,
	},

	"restart": {
		"restart the server",
		"The restart command shuts down the server and restarts",
		&restartServer,
	},

	"encryptdatabase": {
		"encrypt your database",
		"This command encrypts the database containing your bitcoin private keys, identity key, and contracts",
		&cmd.EncryptDatabase{},
	},

	"decryptdatabase": {
		"decrypt your database",
		"This command decrypts the database containing your bitcoin private keys, identity key, and contracts.\n [Warning] doing so may put your bitcoins risk.",
		&cmd.DecryptDatabase{},
	},

	"restore": {
		"restore user data",
		"This command will attempt to restore user data (profile, listings, ratings, etc) by downloading them from the network. This will only work if the IPNS mapping is still available in the DHT. Optionally it will take a mnemonic seed to restore from.",
		&cmd.Restore{},
	},

	"convert": {
		"convert this node to a different coin type",
		"This command will convert the node to use a different crypto-currency",
		&cmd.Convert{},
	},
}
