package storage

import (
	"logaggregator/storage/sqlite"
	"database/sql"
)

type IStorage interface {
	LogRepository() sqlite.LogRepository
}

type storageImpl struct {
	sqlite *sql.DB
}

func NewIStorage(db *sql.DB) IStorage {
	return &storageImpl{
		sqlite: db,
	}
}

func (s *storageImpl) LogRepository() sqlite.LogRepository {
	return sqlite.NewLogRepository(s.sqlite)
}
