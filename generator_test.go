package pkce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	randomString, err := GenerateRandomString(128)
	if assert.Nil(t, err) {
		assert.Equal(t, 128, len(randomString))
	}
}
