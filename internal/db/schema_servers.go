package db

import (
	"database/sql"
	"fmt"
)

func createServersTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS servers (
        id SERIAL PRIMARY KEY,
        hostname TEXT NOT NULL,
        ip_address TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT now()
    );`)
	if err != nil {
		return fmt.Errorf("servers tablosu oluşturulamadı: %w", err)
	}
	return nil
}
