package operations

import (
	"ProjectOlimpos/db/schemas"
	"database/sql"
)

func InitSchema(db *sql.DB) error {
	if err := schemas.EnableTimescale(db); err != nil {
		return err
	}
	if err := schemas.CreateServersTable(db); err != nil {
		return err
	}
	if err := schemas.CreateJobsTable(db); err != nil {
		return err
	}
	if err := schemas.CreateWorkerJoinRequestsTable(db); err != nil {
		return err
	}
	if err := schemas.CreateWorkerConnectionsTable(db); err != nil {
		return err
	}
	if err := schemas.CreateAPIRequestLogsTable(db); err != nil {
		return err
	}
	return nil
}
