package servers

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func DeleteServerByID(db *gorm.DB, id uint) error {
	return db.Delete(&models.Server{}, id).Error
}
