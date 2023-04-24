package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:dreamer@localhost:5432/bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("err while connecting to db", err)
	}

	testQuries = New(con)
	os.Exit(m.Run())
}
