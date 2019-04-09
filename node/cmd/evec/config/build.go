package config

import (
	"github.com/urfave/cli"
)

var (
	// BuildTengo is the state of build command flag "tengo".
	BuildTengo bool

	// BuildEvelyn is the state of build command flag "evelyn".
	BuildEvelyn bool

	// BuildVyper is the state of build command flag "vyper".
	BuildVyper bool

	// BuildSolidity is the state of build command flag "solidity".
	BuildSolidity bool
)

// Ok returns true if all options are consistent.
// Ok returns false and a message if the options are bad.
func Ok(c *cli.Context) (bool, string) {
	if c.NArg() == 0 {
		return false, "no input files"
	}
	if c.IsSet("output") && c.NArg() > 1 {
		return false, "flag --output is used with more then one input file"
	}
	lang := 0
	if BuildTengo {
		lang++
	}
	if BuildEvelyn {
		lang++
	}
	if BuildVyper {
		lang++
	}
	if BuildSolidity {
		lang++
	}
	if lang > 1 {
		return false, "Command error: only one explicit language is allowed"
	}
	return true, ""
}
