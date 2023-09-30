package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}

type PayloadParams struct {
	username string
	duration time.Duration
}

func (p PayloadParams) NewPayload() (payload *Payload, err error) {
	var tokenID uuid.UUID

	tokenID, err = uuid.NewRandom()
	if err != nil {
		return
	}

	payload = &Payload{
		ID:        tokenID,
		Username:  p.username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(p.duration),
	}

	return
}
