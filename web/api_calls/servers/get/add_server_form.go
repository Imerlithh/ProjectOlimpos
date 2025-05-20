package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowAddServerForm(c *gin.Context) {
	c.HTML(http.StatusOK, "layout", gin.H{
		"ActivePage":      "servers",
		"PageTitle":       "Add New Server",
		"Title":           "Add Server",
		"ContentTemplate": "servers_add_content",
	})
}
