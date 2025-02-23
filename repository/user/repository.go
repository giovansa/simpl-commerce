package user

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(dbConn *sql.DB) *Repository {
	return &Repository{
		Db: dbConn,
	}
}
