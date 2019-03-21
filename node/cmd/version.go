package cmd

import (
	"fmt"
	"github.com/evenfound/even-go/node/core"
)

type CoreVersion struct {
	version string
}

func (x *CoreVersion) Execute(args []string) error {

	x.version = core.VERSION

	fmt.Printf("Even Nenwork core version %s", x.version)

	return nil
}
