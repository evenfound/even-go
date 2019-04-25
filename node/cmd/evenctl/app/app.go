package app

import (
	"context"
	"os"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"
	"github.com/evenfound/even-go/node/cmd/evenctl/rpc"
	"github.com/evenfound/even-go/node/cmd/evenctl/tool"
	"github.com/urfave/cli"
)

// Init initializes the application.
func Init() {
}

// Close finalizes the application.
func Close() {
}

// Run starts the application.
func Run() error {

	ctx := context.Background()

	peerCMD, err := rpc.NewPeerCMD(ctx)

	if err != nil {
		return err
	}

	fileCMD, err := rpc.NewFileCMD()

	if err != nil {
		return err
	}

	a := cli.NewApp()
	a.Name = "evenctl"
	a.Usage = "Even Network control tool"
	a.Version = "0.13.0"

	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "show additional information",
		},
	}

	a.Commands = []cli.Command{
		{
			Name:  "call",
			Usage: "Calls a smart contract",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no file name")
				case 1:
					return rpc.Call(c.Args()[0], config.DefaultEntryFunction)
				case 2:
					return rpc.Call(c.Args()[0], c.Args()[1])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "peer.list",
			Usage: "Get peers  list in private network",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return peerCMD.List()
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "peer.send",
			Usage: "Send hash to peer",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no hash name")
				case 1:
					return peerCMD.SendStore(c.Args()[0])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "file.find",
			Usage: "Get file from IPFS by cid",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no file name")
				case 2:
					return fileCMD.GetFileByHash(c.Args()[0], c.Args()[1])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "file.create",
			Usage: "Create Mutable file",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no file name")
				case 2:
					return fileCMD.Create(c.Args()[1], os.Args[2])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "files.mkdir",
			Usage: "Create Directory",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no file name")
				case 1:
					return fileCMD.Mkdir(c.Args()[0])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
		{
			Name:  "file.stat",
			Usage: "Get file stat data",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.Check()
				switch c.NArg() {
				case 0:
					return tool.NewError("no file name")
				case 1:
					return fileCMD.Stat(c.Args()[0])
				}
				return tool.NewError("too many arguments; first argument should be a file name, optional second should be entry function name")
			},
		},
	}

	return a.Run(os.Args)
}
