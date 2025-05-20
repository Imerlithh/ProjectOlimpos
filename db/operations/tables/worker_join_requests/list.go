package worker_join_requests

import (
	"database/sql"
	"fmt"
)

func ListWorkerJoinRequests(db *sql.DB) ([]WorkerJoinRequest, error) {
	rows, err := db.Query(`SELECT id, token, worker_name, ip_address, status, created_at, approved_at FROM worker_join_requests`)
	if err != nil {
		return nil, fmt.Errorf("worker join requests listelenemedi: %w", err)
	}
	defer rows.Close()

	var requests []WorkerJoinRequest
	for rows.Next() {
		var w WorkerJoinRequest
		if err := rows.Scan(&w.ID, &w.Token, &w.WorkerName, &w.IPAddress, &w.Status, &w.CreatedAt, &w.ApprovedAt); err != nil {
			return nil, fmt.Errorf("row scan hatası: %w", err)
		}
		requests = append(requests, w)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration hatası: %w", err)
	}

	return requests, nil
}
