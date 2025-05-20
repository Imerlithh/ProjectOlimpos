package schemas

import (
	"database/sql"
	"fmt"
)

func CreateJobsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS jobs (
        id SERIAL PRIMARY KEY,
        server_id INTEGER NOT NULL REFERENCES servers(id),
        name TEXT NOT NULL,
        status TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT now()
    );`)
	if err != nil {
		return fmt.Errorf("jobs tablosu oluşturulamadı: %w", err)
	}
	return nil
}
