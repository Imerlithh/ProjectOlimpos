package db

import (
	"log"

	"ProjectOlimpos/config"
	"ProjectOlimpos/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectGORM(cfg config.Config) *gorm.DB {
	dsn := config.BuildConnectionString(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("GORM bağlantı hatası: %v", err)
	}

	// Tabloları otomatik oluştur
	if err := db.AutoMigrate(
		&models.Server{},
		&models.APIRequestLog{},
		&models.WorkerJoinRequest{},
		&models.WorkerConnection{},
		&models.Job{},
		&models.TimescaleMetric{},
	); err != nil {
		log.Fatalf("AutoMigrate hatası: %v", err)
	}

	log.Println("GORM ile bağlantı ve şema oluşturma başarılı.")
	DB = db
	return db
}
