package schemas

import (
	"database/sql"
	"fmt"
)

func EnableTimescale(db *sql.DB) error {
	_, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS timescaledb;`)
	if err != nil {
		return fmt.Errorf("Timescale uzantısı yüklenemedi: %w", err)
	}
	return nil
}
