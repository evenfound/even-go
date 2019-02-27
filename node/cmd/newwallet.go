package cmd

import (
	"errors"
	"fmt"
	evenWallet "github.com/evenfound/even-go/node/hdwallet"
	"gopkg.in/AlecAivazis/survey.v1"
)

const (
	PasswordConfirmationError = "Provided passwords is not equal"
	DirectoryError            = "The provided directory does not exists"
	InvalidPhraseError        = "The provided seed phrase is not valid"
)

const (
	AskEnterName = "Enter the name of your wallet : "

	AskEnterPhrase = "Enter the private phrase of your wallet : "

	AskEnterPassword = "Please enter you password : "

	AskEnterPasswordConfirmation = "Please confirm you password : "

	AskEnterCoinType = "Please enter type of coin : "

	AskEnterWalletDataDirectory = "Please data directory [%v][Press ENTER to set default value] :  "
)

type CreateWallet struct {
	evenWallet.HDWallet
}

func (wallet CreateWallet) Execute(args []string) (err error) {

	// asking the user to provide wallet name. This string will be used to create directory ,
	// where will be stored all wallet data
	// if user doesn't  provide the name , the system generates name based on timestamp
	if wallet.WalletName == "" {
		var question = []*survey.Question{
			{
				Name: "name",
				Prompt: &survey.Input{
					Message: AskEnterName,
				},
				Validate: survey.Required,
			},
		}

		var answer = struct {
			Name string
		}{}

		err := survey.Ask(question, &answer)

		if err != nil {
			return err
		}

		wallet.WalletName = answer.Name
	}

	// asking user to provide private seed phrase. The phrase will be used to generate wallet.
	// if the user doesn't provide the phrase,  the system will be generate it randomly
	if wallet.SeedPhrase == "" {
		var question = []*survey.Question{
			{
				Name: "phrase",
				Prompt: &survey.Input{
					Message: AskEnterPhrase,
				},
			},
		}

		var answer = struct {
			Phrase string
		}{}

		err := survey.Ask(question, &answer)

		if err != nil {
			return err
		}

		wallet.SeedPhrase = answer.Phrase
	}

	// Validation seed phrase
	if !evenWallet.ValidatePhrase(wallet.SeedPhrase) {
		return errors.New(InvalidPhraseError)
	}

	// ask the user to provide wallet password  and than confirm it
	// if the provided passwords is not equal than the user receives an error
	if wallet.Password == "" {
		var question = []*survey.Question{
			{
				Name: "password",
				Prompt: &survey.Password{
					Message: AskEnterPassword,
				},
				Validate: survey.Required,
			},
			{
				Name: "confirmation",
				Prompt: &survey.Password{
					Message: AskEnterPasswordConfirmation,
				},
				Validate: survey.Required,
			},
		}

		var answer = struct {
			Password     string
			Confirmation string
		}{}

		err := survey.Ask(question, &answer)

		if err != nil {
			return err
		}

		if answer.Password == answer.Confirmation {
			wallet.Password = answer.Password
		} else {
			return errors.New(PasswordConfirmationError)
		}

	}

	_, err = wallet.Create()

	if err == nil {
		fmt.Println(evenWallet.StatusSuccessfullyCreated)
	}

	return err

}
