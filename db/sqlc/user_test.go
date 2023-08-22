package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tuda4/simple_bank/util"
)

func createRandomUser(t *testing.T) User {
	hashPassword, err := util.HashedPassword(util.RandString(10))
	require.NoError(t, err)
	params := CreateUserParams{
		Username:       util.RandOwner(),
		HashedPassword: hashPassword,
		FullName:       util.RandOwner(),
		Email:          util.RandEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Username, params.Username)
	require.Equal(t, user.HashedPassword, params.HashedPassword)
	require.Equal(t, user.FullName, params.FullName)
	require.Equal(t, user.Email, params.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
}
