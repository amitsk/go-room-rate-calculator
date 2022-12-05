package main

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func main() {
	ctx := context.Background()
	var db *sqlx.DB
	var err error

	// exactly the same as the built-in
	db, err = sqlx.Open("sqlite3", ":memory:")

	// // from a pre-existing sql.DB; note the required driverName
	// db = sqlx.NewDb(sql.Open("sqlite3", ":memory:"), "sqlite3")

	// force a connection and test that it worked
	err = db.Ping()
}
