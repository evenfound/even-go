package rpc

import (
	"regexp"
	"strings"

	"github.com/evenfound/even-go/node/cmd/evenctl/config"
)

func isCorrectFilename(filename string) bool {
	return strings.HasPrefix(filename, config.FilePrefix) ||
		strings.HasPrefix(filename, config.IpfsPrefix)
}

func isCorrectFunction(name string) bool {
	return regexp.MustCompile(`^[^\d\W]\w*$`).MatchString(name)
}
