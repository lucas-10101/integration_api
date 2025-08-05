package db

import (
	"os"

	"github.com/jmoiron/sqlx"
)

type SQLConnection interface{}

const DriverName = "sqlite3"

var SqlConnection *sqlx.DB

func CreateConnection() {

	var err error
	if SqlConnection, err = sqlx.Open(DriverName, os.Getenv("DSN")); err != nil {

	}
}
