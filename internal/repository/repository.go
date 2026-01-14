package repository

import "database/sql"

type PgStorage struct {
	db *sql.DB
}

