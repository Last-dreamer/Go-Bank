package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"bank/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const ()

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot retrieve env var", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBServer)
	if err != nil {
		log.Fatal("err while connecting to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("err while starting server")
	}

}
