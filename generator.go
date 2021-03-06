package pkce

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// LETTERS according to https://tools.ietf.org/html/rfc7636#section-4.1
const LETTERS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-._~"

// GenerateRandomString returns a random string of a given length
func GenerateRandomString(l int) string {
	builder := strings.Builder{}
	builder.Grow(l)

	for i := 0; i < l; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(56))
		builder.WriteString(string(LETTERS[n.Int64()]))
	}
	return builder.String()
}
