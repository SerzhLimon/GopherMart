package repository

import (
	"database/sql"
	"fmt"
)

type Repository interface {
}

type PgStorage struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) (Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("db not init")
	}
	return &PgStorage{
		db: db,
	}, nil
}


