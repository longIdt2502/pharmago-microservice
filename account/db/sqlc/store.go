package db

import "database/sql"

type StoreAccout struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) StoreAccout {
	return StoreAccout{
		db:      db,
		Queries: New(db),
	}
}
