package api_logs

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func ListAPIRequestLogs(db *gorm.DB) ([]models.APIRequestLog, error) {
	var logs []models.APIRequestLog
	err := db.Order("created_at DESC").Find(&logs).Error
	return logs, err
}
