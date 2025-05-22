package api_logs

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func DeleteAPIRequestLogByID(db *gorm.DB, id uint) error {
	return db.Delete(&models.APIRequestLog{}, id).Error
}
