package api

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/tuda4/simple_bank/db/sqlc"
	"github.com/tuda4/simple_bank/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}
