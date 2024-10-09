package cmd

import (
	"database/sql"
	"logaggregator/storage"
)

var SQLite storage.IStorage

func SetDBConnection(database *sql.DB) {
	SQLite = storage.NewIStorage(database)
}
