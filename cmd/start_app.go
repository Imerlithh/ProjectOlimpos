package cmd

import (
	"log"

	"ProjectOlimpos/config"
	"ProjectOlimpos/db/operations"
	"ProjectOlimpos/web"
)

func startApplication() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config yüklenemedi: %v", err)
	}
	config.Current = cfg
	operations.InitDB()
	r := web.SetupRouter()
	r.Run(":8080")
}
