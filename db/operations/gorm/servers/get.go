package servers

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func GetServerByID(db *gorm.DB, id uint) (*models.Server, error) {
	var server models.Server
	err := db.First(&server, id).Error
	return &server, err
}
