// Copyright (c) 2018-2019 The Even Foundation developers
// Use of this source code is governed by an ISC license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/evenfound/even-go/node/core"
)

type CoreVersion struct {
	s string
}

func (ver *CoreVersion) Execute(args []string) error {

	ver.s = core.SprintVersion()

	fmt.Println(ver.s)
	os.Exit(0)

	return nil
}
