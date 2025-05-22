package get

import (
	"net/http"

	"ProjectOlimpos/db"
	gormservers "ProjectOlimpos/db/operations/gorm/servers"
	"github.com/gin-gonic/gin"
)

func ListServersHandler(c *gin.Context) {
	servers, err := gormservers.ListServers(db.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "Sunucular alınamadı: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Servers":         servers,
		"ActivePage":      "servers",
		"PageTitle":       "Sunucular",
		"ContentTemplate": "servers_content",
	})
}
