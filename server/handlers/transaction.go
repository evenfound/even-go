package handlers

import (
	protoApi "github.com/evenfound/even-go/server/api"
	"golang.org/x/net/context"
)

// Transaction is a handler.
type Transaction struct{}

// GetTransactions calls a handler.
func (transaction *Transaction) GetTransactions(_ context.Context, in *protoApi.GetTransactionRequestMessage) (*protoApi.GetTransactionResponseMessage, error) {
	return nil, nil
}
