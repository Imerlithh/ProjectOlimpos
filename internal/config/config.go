package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUrl string
}

func Load() Config {
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		dbUrl = "postgres://postgres:password@localhost:5432/olimpos?sslmode=disable"
		fmt.Println("Uyarı: Ortam değişkeni bulunamadı, varsayılan DB bağlantısı kullanılacak.")
	}
	return Config{DBUrl: dbUrl}
}
