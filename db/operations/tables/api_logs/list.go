package api_logs

import (
	"database/sql"
	"fmt"
)

func ListAPIRequestLogs(db *sql.DB) ([]APIRequestLog, error) {
	rows, err := db.Query(`
        SELECT id, endpoint, method, request_body, response_code, client_ip, created_at
        FROM api_request_logs
        ORDER BY created_at DESC
    `)
	if err != nil {
		return nil, fmt.Errorf("API logs listelenemedi: %w", err)
	}
	defer rows.Close()

	var logs []APIRequestLog
	for rows.Next() {
		var log APIRequestLog
		if err := rows.Scan(&log.ID, &log.Endpoint, &log.Method, &log.RequestBody, &log.ResponseCode, &log.ClientIP, &log.CreatedAt); err != nil {
			return nil, fmt.Errorf("row scan hatası: %w", err)
		}
		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration hatası: %w", err)
	}

	return logs, nil
}
