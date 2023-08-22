package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestHashedPassword(t *testing.T) {
	password := RandString(10)

	hashPass, err := HashedPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashPass)

	err = CheckPasswordHash(password, hashPass)
	require.NoError(t, err)

	wrongPass := RandString(10)
	err = CheckPasswordHash(wrongPass, hashPass)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
