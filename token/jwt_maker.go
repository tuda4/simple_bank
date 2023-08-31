package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secretKey string
}

const secretKeySize = 48

func (jwt JWTMaker) NewJWTMaker() (Maker, error) {
	if len(jwt.secretKey) < secretKeySize {
		return nil, fmt.Errorf("invalid secret key size: %d", secretKeySize)
	}

	return &JWTMaker{jwt.secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (token string, err error) {
	var payload *Payload
	arg := PayloadParams{
		username: username,
		duration: duration,
	}
	payload, err = arg.NewPayload()
	if err != nil {
		return
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err = jwtToken.SignedString([]byte(maker.secretKey))
	return
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}