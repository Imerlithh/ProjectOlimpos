package servers

import (
	"ProjectOlimpos/db/models"
	"gorm.io/gorm"
)

func AddServer(db *gorm.DB, s models.Server) error {
	return db.Create(&s).Error
}
