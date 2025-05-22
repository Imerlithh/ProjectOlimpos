package servers

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func ListServers(db *gorm.DB) ([]models.Server, error) {
	var servers []models.Server
	err := db.Order("created_at DESC").Find(&servers).Error
	return servers, err
}
