package operations

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"database/sql"
	"log"
)

// InitDB initializes GORM and ensures hypertables are created
func InitDB() {
	cfg := config.Load()
	gormDB := db.ConnectGORM(cfg)

	// Hypertable dönüşümünü yap
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Printf("sql.DB erişilemedi: %v", err)
		return
	}

	if err := ensureTimescaleHypertables(sqlDB); err != nil {
		log.Printf("Hypertable kurulumu başarısız: %v", err)
	} else {
		log.Println("Hypertable kontrolü tamamlandı.")
	}
}

func ensureTimescaleHypertables(db *sql.DB) error {
	// timescale_metrics tablosunu hypertable'a çevir
	_, err := db.Exec(`
		SELECT create_hypertable('timescale_metrics', 'timestamp', if_not_exists => TRUE);
	`)
	return err
}
