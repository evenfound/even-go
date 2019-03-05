package hdwallet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"github.com/btcsuite/btcwallet/netparams"
	"github.com/btcsuite/btcwallet/wallet"
	"github.com/evenfound/even-go/node/repo"
	"io"
	"path/filepath"
	"regexp"
	"strings"
)

func GetNetwork(testNet bool) netparams.Params {

	if testNet {
		return netparams.TestNet3Params
	}

	return netparams.MainNetParams
}

func GeneratePath(testnet bool, name string) (string, error) {

	var repoPath, err = repo.GetRepoPath(testnet)

	if err != nil {
		return "", err
	}

	return filepath.Join(repoPath, DefaultWalletDirectory, name), nil

}

func GetLoader(testnet bool, name string) (*wallet.Loader, error) {

	var walletPath, err = GeneratePath(testnet, name)

	if err != nil {
		return nil, err
	}

	var network = GetNetwork(testnet)

	return wallet.NewLoader(network.Params, walletPath, 0255), nil
}

// Creating md5 hash from string
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

//
func Encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

//
func Decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

// Validating seed phrase.
// The seed phrase must contain only letters and spaces
//  The seed phrase must contain [12,24] words
func ValidatePhrase(phrase string) bool {
	if !regexp.MustCompile(`^[a-zA-Z\s]*$`).MatchString(phrase) {
		return false
	}
	arrayWords := strings.Split(phrase, " ")

	if len(arrayWords) < 12 || len(arrayWords) > 24 {
		return false
	}

	return true
}
