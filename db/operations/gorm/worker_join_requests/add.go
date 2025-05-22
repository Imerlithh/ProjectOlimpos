package worker_join_requests

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func AddRequest(db *gorm.DB, r models.WorkerJoinRequest) error {
	return db.Create(&r).Error
}
