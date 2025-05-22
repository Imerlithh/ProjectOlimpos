package main

import (
	"ProjectOlimpos/db/operations"
	"ProjectOlimpos/web"
)

func main() {
	operations.InitDB() // GORM bağlantısı ve schema hazır
	r := web.SetupRouter()
	r.Run(":8080")
}
