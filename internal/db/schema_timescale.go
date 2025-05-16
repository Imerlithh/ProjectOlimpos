package db

import (
	"database/sql"
	"fmt"
)

func enableTimescale(db *sql.DB) error {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS timescaledb;`)
	if err != nil {
		return fmt.Errorf("Timescale uzantısı yüklenemedi: %w", err)
	}
	return nil
}
