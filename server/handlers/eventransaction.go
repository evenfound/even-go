package handlers

import (
	"github.com/evenfound/even-go/server/api"
	"github.com/evenfound/even-go/transaction"
	"golang.org/x/net/context"
)

// EvenTransaction is a handler.
type EvenTransaction struct{}

// Create creates new transaction.
func (t *EvenTransaction) Create(_ context.Context, in *api.EvenTransactionCreateInput) (*api.EvenTransactionResult, error) {
	builder := transaction.NewBuilderOutput(in.Tag, 100)
	builder.
		SetAddress("msnxVmXXoTxVP5DNerVbiJ31AP6snsAmiv").
		SetMessage("30440220790fc006bb8321354456b89e2c6a855e4f77047704b14d799faa73363ee3e182022060b6273f6857154cc26cf7c1b84f79db98b3c014d9869dbbbe2e5c2b57ba06cb").
		SetSource("myPH9LjpgV8TEKCm7pmqRnhLXmd7JPUoCn").
		SetTrunk("QmTfCejgo2wTwqnDJs8Lu1pCNeCrCDuE4GAwkna93zdd7d").
		AddTwig("m111111111111111111111111111111111").
		AddTwig("m222222222222222222222222222222222").
		AddTwig("m333333333333333333333333333333333").
		AddTwig("m444444444444444444444444444444444")
	hash, err := builder.SaveLocal(transaction.FileFormat(in.Format))
	if err != nil {
		return failEvenTransaction(err)
	}
	return successEvenTransaction(string(hash))
}

// Show reads and shows transaction.
func (t *EvenTransaction) Show(_ context.Context, in *api.EvenTransactionInput) (*api.EvenTransactionResult, error) {
	tran, err := transaction.Load(in.Filename)
	if err != nil {
		return failEvenTransaction(err)
	}
	return successEvenTransaction(tran.String())
}

// Analyze analyzes transaction.
func (t *EvenTransaction) Analyze(_ context.Context, in *api.EvenTransactionInput) (*api.EvenTransactionResult, error) {
	return successEvenTransaction("Analyze")
}

// Verify validates transaction.
func (t *EvenTransaction) Verify(_ context.Context, in *api.EvenTransactionInput) (*api.EvenTransactionResult, error) {
	return successEvenTransaction("Verify")
}

func failEvenTransaction(err error) (*api.EvenTransactionResult, error) {
	return &api.EvenTransactionResult{
		Ok:     false,
		Result: err.Error(),
	}, nil
}

func successEvenTransaction(result string) (*api.EvenTransactionResult, error) {
	return &api.EvenTransactionResult{
		Ok:     true,
		Result: result,
	}, nil
}
