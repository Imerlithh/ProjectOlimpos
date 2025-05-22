package cmd

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db/operations"
	"ProjectOlimpos/web"
)

func Run() {
	// Config yoksa /first-run'a yönlendirme middleware'i çalışır.
	if config.IsReady() {
		operations.InitDB()
	}

	r := web.SetupRouter()
	r.Run(":8080")
}
