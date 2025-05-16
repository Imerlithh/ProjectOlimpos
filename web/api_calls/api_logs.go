package api_calls

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListAPIRequestLogsHandler(c *gin.Context) {
	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "DB bağlantı hatası: %v", err)
		return
	}
	defer dbConn.Close()

	logs, err := db.ListAPIRequestLogs(dbConn)
	if err != nil {
		c.String(http.StatusInternalServerError, "Veri çekme hatası: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Logs":       logs,
		"ActivePage": "api-logs",
		"PageTitle":  "API Logs",
		"Title":      "API Logs",
	})
}
