package schemas

import (
	"database/sql"
	"fmt"
)

func CreateAPIRequestLogsTable(db *sql.DB) error {
	_, err := db.Exec(`
    CREATE TABLE IF NOT EXISTS api_request_logs (
        id SERIAL PRIMARY KEY,
        endpoint TEXT NOT NULL,
        method TEXT NOT NULL,
        request_body TEXT,
        response_code INTEGER NOT NULL,
        client_ip TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT now()
    );`)
	if err != nil {
		return fmt.Errorf("api_request_logs tablosu oluşturulamadı: %w", err)
	}
	return nil
}
