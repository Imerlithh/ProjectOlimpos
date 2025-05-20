package main

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"ProjectOlimpos/db/operations"
	"ProjectOlimpos/web"
	"log"
)

func main() {
	InitDB()
	r := web.SetupRouter()
	r.Run(":8080")
}

func InitDB() {
	cfg := config.Load()

	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer dbConn.Close()

	log.Println("Veritabanına başarıyla bağlandı.")

	if err := operations.InitSchema(dbConn); err != nil {
		log.Fatalf("Şema oluşturulamadı: %v", err)
	}

	log.Println("Şema başarıyla oluşturuldu.")
}
