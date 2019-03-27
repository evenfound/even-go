package app

import (
	"github.com/evenfound/even-go/node/cmd/evec/config"
	"github.com/evenfound/even-go/node/cmd/evec/tool"
	"os"

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
	a.Name = "evec"
	a.Usage = "Even Smart Contract compiler"
	a.Version = "0.0.1"

	a.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "show additional information",
		},
	}

	a.Commands = []cli.Command{
		{
			Name:  "clean",
			Usage: "remove object files and cached files",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				if ok, msg := config.Ok(); !ok {
					return tool.NewError(msg)
				}
				return clean()
			},
		},
		{
			Name:  "build",
			Usage: "compile program(s)",
			Action: func(c *cli.Context) error {
				config.Debug = c.GlobalBool("debug")
				config.BuildTengo = c.Bool("tengo")
				config.BuildEvelyn = c.Bool("evelyn")
				config.BuildVyper = c.Bool("vyper")
				config.BuildSolidity = c.Bool("solidity")
				if ok, msg := config.Ok(); !ok {
					return tool.NewError(msg)
				}
				if c.NArg() == 0 {
					return buildWorkDir()
				}
				return buildFiles(c.Args())
			},
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "tengo",
					Usage: "force compile Tengo sources",
				},
				cli.BoolFlag{
					Name:  "evelyn",
					Usage: "force compile Evelyn sources",
				},
				cli.BoolFlag{
					Name:  "vyper",
					Usage: "force compile Vyper sources",
				},
				cli.BoolFlag{
					Name:  "solidity",
					Usage: "force compile Solidity sources",
				},
				cli.StringFlag{
					Name:  "output,o",
					Usage: "name of output binary file",
				},
			},
		},
	}

	return a.Run(os.Args)
}
