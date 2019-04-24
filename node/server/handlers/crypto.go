package handlers

import (
	"fmt"

	"github.com/evenfound/even-go/node/crypto"
	"github.com/evenfound/even-go/node/server/api"
	"golang.org/x/net/context"
)

// Crypto is a handler.
type Crypto struct{}

// Sign signs a message and returns the signature.
func (c *Crypto) Sign(ctx context.Context, in *api.SignInput) (*api.SignResult, error) {
	signature, err := crypto.Sign(in.Message, in.Privkey)
	if err != nil {
		return fail(err)
	}
	return success(signature)
}

// Verify recovers the account which was used to sign a message.
func (c *Crypto) Verify(ctx context.Context, in *api.VerifyInput) (*api.SignResult, error) {
	valid, err := crypto.Verify(in.Message, in.Signature, in.Pubkey)
	if err != nil {
		return fail(err)
	}
	return success(fmt.Sprint(valid))
}

func fail(err error) (*api.SignResult, error) {
	return &api.SignResult{
		Ok:     false,
		Result: err.Error(),
	}, nil
}

func success(result string) (*api.SignResult, error) {
	return &api.SignResult{
		Ok:     true,
		Result: result,
	}, nil
}
