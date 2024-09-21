package product

import "database/sql"

type Repository struct {
	Db *sql.DB
}

func NewRepository(dbConn *sql.DB) *Repository {
	return &Repository{
		Db: dbConn,
	}
}
