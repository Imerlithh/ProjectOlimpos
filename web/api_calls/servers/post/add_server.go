package post

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"ProjectOlimpos/db/operations/tables/servers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddServerHandler(c *gin.Context) {
	hostname := c.PostForm("hostname")
	ipAddress := c.PostForm("ip_address")

	if hostname == "" || ipAddress == "" {
		c.String(http.StatusBadRequest, "Hostname ve IP Address gerekli.")
		return
	}

	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "DB bağlantı hatası: %v", err)
		return
	}
	defer dbConn.Close()

	if err := servers.AddServer(dbConn, hostname, ipAddress); err != nil {
		c.String(http.StatusInternalServerError, "Server ekleme hatası: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/servers")
}
