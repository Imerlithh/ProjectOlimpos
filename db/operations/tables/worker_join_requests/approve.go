package worker_join_requests

import (
	"database/sql"
	"fmt"
)

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
