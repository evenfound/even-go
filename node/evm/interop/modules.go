package interop

import (
	"github.com/d5/tengo/objects"
	"github.com/d5/tengo/script"
	"github.com/d5/tengo/stdlib"
)

// addBuiltinModules creates the module map and populates it with modules and functions.
func addBuiltinModules(env *Environment, s *script.Script) {
	mods := objects.NewModuleMap()

	// Script does not include any modules by default
	mods.AddBuiltinModule("times", stdlib.BuiltinModules["times"])

	// Custom module even
	mods.AddBuiltinModule("even", map[string]objects.Object{
		"println": &objects.UserFunction{Name: "println", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			msg, _ := objects.ToString(args[0])
			return &objects.Int{Value: int64(env.evenPrintln(msg))}, nil
		}},
		"addString": &objects.UserFunction{Name: "addString", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			str, _ := objects.ToString(args[0])
			return &objects.Int{Value: int64(env.addString(str))}, nil
		}},
		"hash": &objects.UserFunction{Name: "hash", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			msg, _ := objects.ToString(args[0])
			return &objects.String{Value: env.evenHash(msg)}, nil
		}},
		"sign": &objects.UserFunction{Name: "sign", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			msg, _ := objects.ToString(args[0])
			privkey, _ := objects.ToString(args[1])
			signature, err := env.evenSign(msg, privkey)
			if err != nil {
				return nil, err
			}
			return &objects.String{Value: signature}, nil
		}},
		"verify": &objects.UserFunction{Name: "verify", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			msg, _ := objects.ToString(args[0])
			signature, _ := objects.ToString(args[1])
			pubkey, _ := objects.ToString(args[2])
			valid, err := env.evenVerify(msg, signature, pubkey)
			if err != nil {
				return nil, err
			}
			if valid {
				return objects.TrueValue, nil
			}
			return objects.FalseValue, nil
		}},
		"createWallet": &objects.UserFunction{Name: "createWallet", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			name, _ := objects.ToString(args[0])
			seed, _ := objects.ToString(args[1])
			h, err := env.evenCreateWallet(name, seed)
			if err != nil {
				return nil, err
			}
			return &objects.Int{Value: int64(h)}, nil
		}},
	})

	// Custom module wallet
	mods.AddBuiltinModule("wallet", map[string]objects.Object{
		"save": &objects.UserFunction{Name: "save", Value: func(args ...objects.Object) (ret objects.Object, err error) {
			w, _ := objects.ToInt64(args[0])
			password, _ := objects.ToString(args[1])
			err = env.walletSave(handle(w), password)
			if err != nil {
				ret = objects.FalseValue
				return
			}
			ret = objects.TrueValue
			return
		}},
	})

	s.SetImports(mods)
}
