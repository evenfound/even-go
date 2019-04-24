package crypto

import (
	"golang.org/x/crypto/sha3"
)

// Hash calculates 256-bit hash of message.
func Hash(message []byte) []byte {
	digest := sha3.Sum256(message)
	return digest[:]
}
