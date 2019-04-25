package handlers

import (
	"fmt"

	"github.com/evenfound/even-go/crypto"
	"github.com/evenfound/even-go/server/api"
	"golang.org/x/net/context"
)

// Crypto is a handler.
type Crypto struct{}

// Sign signs a message and returns the signature.
func (c *Crypto) Sign(_ context.Context, in *api.SignInput) (*api.SignResult, error) {
	signature, err := crypto.Sign(in.Message, in.Privkey)
	if err != nil {
		return failSign(err)
	}
	return successSign(signature)
}

// Verify recovers the account which was used to sign a message.
func (c *Crypto) Verify(_ context.Context, in *api.VerifyInput) (*api.SignResult, error) {
	valid, err := crypto.Verify(in.Message, in.Signature, in.Pubkey)
	if err != nil {
		return failSign(err)
	}
	return successSign(fmt.Sprint(valid))
}

func failSign(err error) (*api.SignResult, error) {
	return &api.SignResult{
		Ok:     false,
		Result: err.Error(),
	}, nil
}

func successSign(result string) (*api.SignResult, error) {
	return &api.SignResult{
		Ok:     true,
		Result: result,
	}, nil
}
