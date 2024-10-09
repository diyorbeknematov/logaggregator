package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectToSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "/home/diyorbek/Desktop/golang/github.com/imtihon/LogAggregator/logaggregator/database/logs.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
