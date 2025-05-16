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
		dbUrl = "postgres://olimpos:supersecret@oduncutest1.cc.itu.edu.tr:5432/olimposdb?sslmode=disable"
		fmt.Println("Uyarı: Ortam değişkeni bulunamadı, varsayılan DB bağlantısı kullanılacak.")
	}
	return Config{DBUrl: dbUrl}
}
