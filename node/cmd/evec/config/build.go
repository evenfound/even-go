package config

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
func Ok() (bool, string) {
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
