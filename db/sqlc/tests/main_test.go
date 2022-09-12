package db

import (
	"database/sql"
	"log"
	"os"
	db "tara/simplebank/db/sqlc"
	"tara/simplebank/db/util"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config,err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
