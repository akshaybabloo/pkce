// Proof Key for Code Exchange (PKCE) helps public clients to get access tokens and refresh tokens. PKCE is based on
// RFC7636, more on that can be found at https://tools.ietf.org/html/rfc7636 and https://auth0.com/docs/flows/authorization-code-flow-with-proof-key-for-code-exchange-pkce.
package pkce

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

type Pkce struct {

	// Optional, RandomString, this will be converted to Pkce.ChallengeCode
	RandomString string

	// Optional, Length of the Pkce.VerifyCode. When RandomString is provided, the Length is ignored.
	// The value should be minimum 43 and maximum 128.
	Length int
}

// When Pkce.RandomString is provided, VerifyCode is similar to it
// if not, VerifyCode returns a random string based on Pkce.Length.
func (p *Pkce) VerifyCode() (string, error) {
	if p.RandomString == "" {

		if p.Length == 0 {
			return "", errors.New("length should be greater than 0")
		}

		if p.Length < 43 || p.Length > 128 {
			return "", errors.New("length should be >=43 and <=128")
		}

		randomString, err := GenerateRandomString(p.Length)
		if err != nil {
			return "", err
		}
		p.RandomString = randomString
		return randomString, nil
	} else {

		if len(p.RandomString) < 43 || len(p.RandomString) > 128 {
			return "", errors.New("length should be >=43 and <=128")
		}

		return p.RandomString, nil
	}
}

// ChallengeCode returns a challenge code as mentioned in https://tools.ietf.org/html/rfc7636#section-4.2.
// The code is based on
// 	code_challenge = BASE64URL-ENCODE(SHA256(ASCII(code_verifier)))
func (p *Pkce) ChallengeCode() string {
	hash := sha256.Sum256([]byte(p.RandomString))
	toBase64 := base64.RawURLEncoding.EncodeToString(hash[:])
	return toBase64
}
