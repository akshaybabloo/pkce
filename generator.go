package pkce

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// LETTERS according to https://tools.ietf.org/html/rfc7636#section-4.1
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"

// GenerateRandomString returns a random string of a given length
func GenerateRandomString(l int) (string, error) {
	builder := strings.Builder{}
	builder.Grow(l)

	for i := 0; i < l; i++ {
		number, err := generateRandomNumber()
		if err != nil {
			return "", err
		}
		builder.WriteString(string(LETTERS[number]))
	}
	return builder.String(), nil
}

func generateRandomNumber() (int64, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(56))
	if err != nil {
		return 0, err
	}
	return n.Int64(), nil
}
