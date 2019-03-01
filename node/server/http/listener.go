package http

import (
	"encoding/json"
	"fmt"
	"github.com/evenfound/even-go/node/hdwallet"
	"net/http"
)

type NewAccountResponse struct {
	Account string `json:"account"`
	Coin    string `json:"coin"`
	Id      uint32 `json:"id"`
}

type AccountResponse struct {
	Name string `json:"name"`
	Id   uint32 `json:"id"`
}

type ListAccountResponse []AccountResponse


// Creating a new wallet
func NewWallet(w http.ResponseWriter, r *http.Request) *ErrorResponse {

	decoder := json.NewDecoder(r.Body)

	var data hdwallet.HDWallet

	err := decoder.Decode(&data)

	if err != nil {
		fmt.Println(err)
	}

	_, err = data.Create()

	if err != nil {
		return &ErrorResponse{
			Message: err.Error(),
			Code:    400,
		}
	}

	var response = successResponse{}

	jsoned, err := json.Marshal(&response)

	if err != nil {
		return &ErrorResponse{
			Message: err.Error(),
			Code:    400,
		}
	}

	w.Write(jsoned)

	return nil
}

// Creating a new account
func NewAccount(w http.ResponseWriter, r *http.Request) *ErrorResponse {

	decoder := json.NewDecoder(r.Body)

	var data hdwallet.AccountManager

	err := decoder.Decode(&data)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	wallet, err := data.Authorize()

	if err != nil {
		return generateError(err.Error(), 0)
	}

	accountID, err := data.NewAccount(*wallet)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	jsoned, err := json.Marshal(NewAccountResponse{
		Account: data.AccountName,
		Coin:    hdwallet.AvailableCoinTypes[data.Coin],
		Id:      accountID,
	})

	if err != nil {
		return generateError(err.Error(), 0)
	}

	w.Write(jsoned)

	return nil
}

// Getting list accounts of wallet by coin
func ListAccounts(w http.ResponseWriter, r *http.Request) *ErrorResponse {



	decoder := json.NewDecoder(r.Body)

	var data hdwallet.WalletAuth

	err := decoder.Decode(&data)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	wallet, err := data.Authorize()

	if err != nil {
		return generateError(err.Error(), http.StatusUnauthorized)
	}

	var scope = data.GetKeyScope()

	accounts, err := wallet.Accounts(scope)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	var response = ListAccountResponse{}

	for _, raw := range accounts.Accounts {
		response = append(response, AccountResponse{
			Name: raw.AccountName,
			Id:   raw.AccountNumber,
		})
	}

	jsoned, err := json.Marshal(response)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	w.Write(jsoned)

	return nil
}

// Generating a new address
func NewAddress(w http.ResponseWriter, r *http.Request) *ErrorResponse {

	decoder := json.NewDecoder(r.Body)

	var data hdwallet.AddressManager

	err := decoder.Decode(&data)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	wallet, err := data.Authorize()

	if err != nil {
		return generateError(err.Error(), 0)
	}

	data.SetWallet(wallet)

	addresses, err := data.NewAddress()

	if err != nil {
		return generateError(err.Error(), 0)
	}

	jsoned, err := json.Marshal(addresses)

	if err != nil {
		return generateError(err.Error(), 0)
	}

	w.Write(jsoned)

	return nil
}

func generateError(message string, code int) *ErrorResponse {

	if code == 0 {
		code = http.StatusBadRequest
	}

	return &ErrorResponse{
		Message: message,
		Code:    code,
	}
}
