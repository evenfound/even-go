package tool

import (
	"fmt"
	"reflect"
)

// TR (trace) prints it's arguments and the types of arguments.
func TR(title string, aa ...interface{}) {
	fmt.Print(title, " ")
	for _, a := range aa {
		fmt.Print(a, " (", reflect.TypeOf(a), ") ")
	}
	fmt.Println()
}
