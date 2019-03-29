package app

import (
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
	a := cli.NewApp()
	a.Name = "evenctl"
	a.Usage = "Even Network control tool"
	a.Version = "0.0.1"

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
	}

	return a.Run(os.Args)
}
