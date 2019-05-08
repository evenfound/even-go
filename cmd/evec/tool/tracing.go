package tool

import (
	"fmt"
	"reflect"
)

// TR (trace) prints it's arguments and the types of arguments.
func TR(prefix string, aa ...interface{}) {
	fmt.Print(prefix)
	for _, a := range aa {
		fmt.Print(a, " (", reflect.TypeOf(a), ") ")
	}
	fmt.Println()
}
