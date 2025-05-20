package servers

import (
	"database/sql"
	"fmt"
)

func ListServers(db *sql.DB) ([]Server, error) {
	rows, err := db.Query(`SELECT id, hostname, ip_address, created_at FROM servers`)
	if err != nil {
		return nil, fmt.Errorf("servers listelenemedi: %w", err)
	}
	defer rows.Close()

	var servers []Server
	for rows.Next() {
		var s Server
		if err := rows.Scan(&s.ID, &s.Hostname, &s.IPAddress, &s.CreatedAt); err != nil {
			return nil, fmt.Errorf("row scan hatası: %w", err)
		}
		servers = append(servers, s)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration hatası: %w", err)
	}

	return servers, nil
}
