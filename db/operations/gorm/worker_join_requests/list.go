package worker_join_requests

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func ListRequests(db *gorm.DB) ([]models.WorkerJoinRequest, error) {
	var requests []models.WorkerJoinRequest
	err := db.Order("created_at DESC").Find(&requests).Error
	return requests, err
}
