package api_calls

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
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

	servers, err := db.ListServers(dbConn)
	if err != nil {
		c.String(http.StatusInternalServerError, "Veri çekme hatası: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Servers":    servers,
		"ActivePage": "servers",
	})
}

func ShowAddServerForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add_server.tmpl", gin.H{
		"ActivePage": "servers",
	})
}

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

	if err := db.AddServer(dbConn, hostname, ipAddress); err != nil {
		c.String(http.StatusInternalServerError, "Server ekleme hatası: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/servers")
}
