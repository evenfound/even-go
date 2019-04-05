package interop

import "github.com/d5/tengo/script"

// addBuiltinVariables adds global script variables accessible from a Go program.
func addBuiltinVariables(s *script.Script) (err error) {
	add := func(name string) {
		if err != nil {
			return
		}
		err = s.Add(name, "")
	}

	add("result")

	return
}
