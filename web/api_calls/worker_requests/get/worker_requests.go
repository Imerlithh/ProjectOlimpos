package get

import (
	"ProjectOlimpos/db"
	gormwjr "ProjectOlimpos/db/operations/gorm/worker_join_requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListWorkerJoinRequestsHandler(c *gin.Context) {
	requests, err := gormwjr.ListRequests(db.DB)
	if err != nil {
		c.String(http.StatusInternalServerError, "İstekler alınamadı: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Requests":        requests,
		"ActivePage":      "worker-requests",
		"PageTitle":       "Worker Join Requests",
		"ContentTemplate": "worker_requests_content",
	})
}
