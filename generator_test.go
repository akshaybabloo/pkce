package pkce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomString(t *testing.T) {
	randomString := GenerateRandomString(128)
	assert.Equal(t, 128, len(randomString))
}
