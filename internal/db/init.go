package db

import "database/sql"

func InitSchema(db *sql.DB) error {
	if err := enableTimescale(db); err != nil {
		return err
	}
	if err := createServersTable(db); err != nil {
		return err
	}
	if err := createJobsTable(db); err != nil {
		return err
	}
	if err := createWorkerJoinRequestsTable(db); err != nil {
		return err
	}
	if err := createWorkerConnectionsTable(db); err != nil {
		return err
	}
	if err := createAPIRequestLogsTable(db); err != nil {
		return err
	}
	return nil
}
