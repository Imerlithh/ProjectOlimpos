package worker_join_requests

import "database/sql"

type WorkerJoinRequest struct {
	ID         int
	Token      string
	WorkerName string
	IPAddress  string
	Status     string
	CreatedAt  string
	ApprovedAt sql.NullString
}
