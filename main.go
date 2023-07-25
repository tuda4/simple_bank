package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tuda4/simple_bank/api"
	db "github.com/tuda4/simple_bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	connect, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(connect)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
