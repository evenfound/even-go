package hdwallet

import (
	"fmt"
	"reflect"
)

// must be error-free, panic otherwise.
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// tr (trace) prints it's arguments and the types of arguments.
func tr(prefix string, aa ...interface{}) {
	fmt.Print(prefix)
	for _, a := range aa {
		fmt.Print(a, " (", reflect.TypeOf(a), ") ")
	}
	fmt.Println()
}
