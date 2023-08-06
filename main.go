package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tuda4/simple_bank/api"
	db "github.com/tuda4/simple_bank/db/sqlc"
	"github.com/tuda4/simple_bank/util"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:123456@localhost:5432/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8080"
// )

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config::::", err)
	}

	connect, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	store := db.NewStore(connect)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
