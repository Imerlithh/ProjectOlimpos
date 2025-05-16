package db

import (
	"database/sql"
	"fmt"
)

func createWorkerConnectionsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS worker_connections (
        id SERIAL PRIMARY KEY,
        worker_id INTEGER NOT NULL REFERENCES servers(id),
        ip_address TEXT NOT NULL,
        connected_at TIMESTAMPTZ DEFAULT now(),
        disconnected_at TIMESTAMPTZ,
        status TEXT NOT NULL DEFAULT 'active'
    );`)
	if err != nil {
		return fmt.Errorf("worker_connections tablosu oluşturulamadı: %w", err)
	}
	return nil
}
