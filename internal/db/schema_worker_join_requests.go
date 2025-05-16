package db

import (
	"database/sql"
	"fmt"
)

func createWorkerJoinRequestsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS worker_join_requests (
        id SERIAL PRIMARY KEY,
        token TEXT NOT NULL UNIQUE,
        worker_name TEXT,
        ip_address TEXT,
        status TEXT NOT NULL DEFAULT 'pending',
        created_at TIMESTAMPTZ DEFAULT now(),
        approved_at TIMESTAMPTZ
    );`)
	if err != nil {
		return fmt.Errorf("worker_join_requests tablosu oluşturulamadı: %w", err)
	}
	return nil
}
