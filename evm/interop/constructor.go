package interop

import "github.com/d5/tengo/script"

// NewEnvironment creates new smart contract environment.
func NewEnvironment(src []byte) (*Environment, error) {
	s := script.New(src)
	if err := addBuiltinVariables(s); err != nil {
		return nil, err
	}

	env := Environment{}
	addBuiltinModules(&env, s)

	runner, err := s.Compile()
	if err != nil {
		return nil, err
	}

	env.runner = runner
	env.wallets = make([]wallet, 0)
	env.strings = make([]string, 0)

	return &env, nil
}
