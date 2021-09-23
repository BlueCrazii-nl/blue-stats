package utils

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDatabase(dbPath string) (*sql.DB, error) {
	return sql.Open("sqlite3", dbPath)
}
