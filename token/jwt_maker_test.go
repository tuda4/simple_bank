package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"github.com/tuda4/simple_bank/util"
)

func TestJWTMaker(t *testing.T) {
	jwt := JWTMaker{
		secretKey: util.RandString(48),
	}
	maker, err := jwt.NewJWTMaker()
	require.NoError(t, err)

	username := util.RandOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := time.Now().Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, payload.Username, username)
	require.WithinDuration(t, payload.ExpiredAt, expiredAt, time.Second)
	require.WithinDuration(t, payload.IssuedAt, issuedAt, time.Second)

}

func TestExpiredJWTToken(t *testing.T) {
	jwt := JWTMaker{
		util.RandString(48),
	}
	maker, err := jwt.NewJWTMaker()
	require.NoError(t, err)
	require.NotEmpty(t, maker)

	token, err := maker.CreateToken(util.RandOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	p := PayloadParams{
		username: util.RandOwner(),
		duration: time.Minute,
	}
	payload, err := p.NewPayload()
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)

	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := JWTMaker{secretKey: util.RandString(48)}.NewJWTMaker()
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
