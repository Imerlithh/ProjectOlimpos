package servers

import (
	"database/sql"
	"fmt"
)

func AddServer(db *sql.DB, hostname, ipAddress string) error {
	_, err := db.Exec(`
        INSERT INTO servers (hostname, ip_address)
        VALUES ($1, $2)
    `, hostname, ipAddress)
	if err != nil {
		return fmt.Errorf("server eklenemedi: %w", err)
	}
	return nil
}
