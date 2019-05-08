package config

import "fmt"

var (
	// Debug is the state of global flag "debug".
	Debug bool
)

// Show prints the current configuration.
func Show() {
	fmt.Println("OK")
}

// Check makes additional checks
func Check() {
}
