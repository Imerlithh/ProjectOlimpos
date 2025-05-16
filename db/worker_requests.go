package db

import (
	"database/sql"
	"fmt"
)

type WorkerJoinRequest struct {
	ID         int
	Token      string
	WorkerName string
	IPAddress  string
	Status     string
	CreatedAt  string
	ApprovedAt sql.NullString
}

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

func ApproveWorkerJoinRequest(db *sql.DB, id int) error {
	_, err := db.Exec(`
        UPDATE worker_join_requests
        SET status = 'approved', approved_at = now()
        WHERE id = $1
    `, id)
	if err != nil {
		return fmt.Errorf("worker request onaylanamadı: %w", err)
	}
	return nil
}
