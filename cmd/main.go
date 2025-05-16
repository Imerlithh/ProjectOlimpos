package main

import (
	"ProjectOlimpos/internal/config"
	"ProjectOlimpos/internal/db"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer dbConn.Close()

	fmt.Println("Veritabanına başarıyla bağlandı.")

	err = db.InitSchema(dbConn)
	if err != nil {
		log.Fatalf("Şema oluşturulamadı: %v", err)
	}

	fmt.Println("Şema başarıyla oluşturuldu.")
}
