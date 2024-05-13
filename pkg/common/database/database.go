package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import PostgreSQL driver
)

var DB *sqlx.DB

func ConnectDB(dsn string) error {
	// Open a connection to the database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
