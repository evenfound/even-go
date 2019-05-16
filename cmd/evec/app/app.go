package app

import (
	"os"

	"github.com/evenfound/even-go/cmd/evec/config"
	"github.com/jawher/mow.cli"
	"github.com/pkg/errors"
)

// Init initializes the application.
func Init() {
}

// Close finalizes the application.
func Close() {
}

// Run starts the application.
func Run() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	a := cli.App("evec", "Even Smart Contract compiler.")

	a.BoolOptPtr(&config.Debug, "d debug", false, "show additional information")

	a.Command("clean", "remove intermediate files", cmdClean)
	a.Command("build", "compile program(s)", cmdBuild)

	return a.Run(os.Args)
}

func cmdClean(c *cli.Cmd) {
	c.Action = func() {
		if ok, msg := config.Ok(); !ok {
			panic(errors.New(msg))
		}
		if err := clean(); err != nil {
			panic(err)
		}
	}
}

func cmdBuild(c *cli.Cmd) {
	c.BoolOptPtr(&config.BuildTengo, "t tengo", false, "force compile Tengo sources")
	c.BoolOptPtr(&config.BuildVyper, "v vyper", false, "force compile Vyper sources")
	c.BoolOptPtr(&config.BuildEvelyn, "e evelyn", false, "force compile Evelyn sources")
	output := c.StringOpt("o output", "", "name of the output binary file; or 'ipfs' to store in the IPFS")
	c.Spec = "FILE... [--tengo] [--vyper] [--evelyn] [--output ...]"
	files := c.StringsArg("FILE", nil, "file names to build")
	c.Action = func() {
		if ok, msg := config.Ok(); !ok {
			panic(errors.New(msg))
		}
		if err := buildFiles(*files, *output); err != nil {
			panic(err)
		}
	}
}
