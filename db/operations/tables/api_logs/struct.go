package api_logs

import (
	"database/sql"
)

type APIRequestLog struct {
	ID           int
	Endpoint     string
	Method       string
	RequestBody  sql.NullString
	ResponseCode int
	ClientIP     string
	CreatedAt    string
}
