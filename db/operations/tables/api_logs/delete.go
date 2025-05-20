package api_logs

import (
	"database/sql"
	"fmt"
)

func DeleteAPIRequestLogByID(db *sql.DB, id int) error {
	query := `DELETE FROM api_request_logs WHERE id = $1`
	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("silme hatası: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("silinen satır sayısı alınamadı: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("id %d için kayıt bulunamadı", id)
	}

	return nil
}
