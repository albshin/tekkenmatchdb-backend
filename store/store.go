package store

import (
	"github.com/jmoiron/sqlx"
)

type PGStore struct {
	*sqlx.DB
}

func Open(dsn string) (*PGStore, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &PGStore{db}, nil
}

func (db *PGStore) Close() error {
	return db.Close()
}
