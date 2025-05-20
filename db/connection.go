package db

import (
	"ProjectOlimpos/config"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	dsn := config.BuildConnectionString(cfg)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// Basit bağlantı kontrolü
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Veritabanına başarıyla bağlandı.")
	return db, nil
}
