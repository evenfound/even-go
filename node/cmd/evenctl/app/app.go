package app

import (
	"os"

	"github.com/evenfound/even-go/node/cmd/evenctl/tool"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"

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
				if c.NArg() == 0 {
					return tool.NewError("no file name")
				}
				if c.NArg() > 1 {
					return tool.NewError("too many file names (only one allowed)")
				}
				return call(c.Args()[0])
			},
		},
	}

	return a.Run(os.Args)
}
