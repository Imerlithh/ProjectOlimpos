package db

import (
	"ProjectOlimpos/internal/config"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	return sql.Open("pgx", cfg.DBUrl)
}
