package post

import (
	"net/http"

	"ProjectOlimpos/db"
	"ProjectOlimpos/db/models"
	gormservers "ProjectOlimpos/db/operations/gorm/servers"
	"github.com/gin-gonic/gin"
)

func AddServerHandler(c *gin.Context) {
	hostname := c.PostForm("hostname")
	ip := c.PostForm("ip_address")

	if hostname == "" || ip == "" {
		c.String(http.StatusBadRequest, "Hostname ve IP zorunludur")
		return
	}

	server := models.Server{
		Hostname:  hostname,
		IPAddress: ip,
	}

	if err := gormservers.AddServer(db.DB, server); err != nil {
		c.String(http.StatusInternalServerError, "Sunucu eklenemedi: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/servers")
}
