package get

import (
	"net/http"

	"ProjectOlimpos/db"
	gormlogs "ProjectOlimpos/db/operations/gorm/api_logs"
	"github.com/gin-gonic/gin"
)

func ListAPIRequestLogsHandler(c *gin.Context) {
	logs, err := gormlogs.ListAPIRequestLogs(db.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "Veriler alınamadı: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Logs":            logs,
		"ActivePage":      "api-logs",
		"PageTitle":       "API Logs",
		"ContentTemplate": "api_logs_content",
	})
}
