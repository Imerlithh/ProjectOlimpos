package db

import (
	"../config"
	"database/sql"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	return sql.Open("pgx", cfg.DBUrl)
}
