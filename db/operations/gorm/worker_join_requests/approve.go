package worker_join_requests

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func UpdateRequestStatus(db *gorm.DB, id uint, status string) error {
	return db.Model(&models.WorkerJoinRequest{}).
		Where("id = ?", id).
		Update("status", status).Error
}
