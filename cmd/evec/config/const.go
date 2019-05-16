package config

import (
	"strings"
)

const (
	// WorkDir is current working directory.
	WorkDir = "."

	// TengoExt is a Tengo source code file name extension.
	TengoExt = ".tgo"

	// EvelynExt is a Evelyn source code file name extension.
	EvelynExt = ".evl"

	// VyperExt is a Vyper source code file name extension.
	VyperExt = ".vy"

	// CompiledExt is a compiled program file name extension.
	CompiledExt = ".out"
)

// LooksLikeSourceFile returns true if the filename can represent a source file.
func LooksLikeSourceFile(filename string) bool {
	return strings.HasSuffix(filename, TengoExt) ||
		strings.HasSuffix(filename, EvelynExt) ||
		strings.HasSuffix(filename, VyperExt)
}
