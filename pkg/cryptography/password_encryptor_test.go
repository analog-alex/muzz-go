package cryptography

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)

	assert.Nil(t, err)
	assert.NotEmpty(t, hash)
}

func TestCheckPasswordHash_Success(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)

	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	isValid := CheckPasswordHash(hash, password)

	assert.True(t, isValid)
}

func TestCheckPasswordHash_Fail(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)

	assert.Nil(t, err)
	assert.NotEmpty(t, hash)

	isValid := CheckPasswordHash(hash, "invalid_password")

	assert.False(t, isValid)
}
