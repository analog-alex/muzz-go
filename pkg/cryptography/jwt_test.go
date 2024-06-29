package cryptography

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateJWToken(t *testing.T) {
	token, err := GenerateJWToken("123")

	assert.Nil(t, err)
	assert.NotNil(t, token)

	assert.Less(t, 16, len(token))
}

func TestVerifyJWToken_Success(t *testing.T) {
	token, err := GenerateJWToken("123")

	assert.Nil(t, err)
	assert.NotNil(t, token)

	userId, err := VerifyJWToken(token)

	assert.Nil(t, err)
	assert.Equal(t, "123", userId)
}

func TestVerifyJWToken_Fail(t *testing.T) {
	token := "invalid_token"

	userId, err := VerifyJWToken(token)

	assert.NotNil(t, err)
	assert.Equal(t, "", userId)
}
