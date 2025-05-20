package get

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"ProjectOlimpos/db/operations/tables/servers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListServersHandler(c *gin.Context) {
	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "DB bağlantı hatası: %v", err)
		return
	}
	defer dbConn.Close()

	servers, err := servers.ListServers(dbConn)
	if err != nil {
		c.String(http.StatusInternalServerError, "Veri çekme hatası: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"ActivePage":      "servers",
		"Servers":         servers,
		"PageTitle":       "Servers",
		"ContentTemplate": "servers_content",
	})
}
