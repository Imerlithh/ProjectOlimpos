package get

import (
	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"ProjectOlimpos/db/operations/tables/worker_join_requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListWorkerRequestsHandler(c *gin.Context) {
	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "DB bağlantı hatası: %v", err)
		return
	}
	defer dbConn.Close()

	requests, err := worker_join_requests.ListWorkerJoinRequests(dbConn)
	if err != nil {
		c.String(http.StatusInternalServerError, "Veri çekme hatası: %v", err)
		return
	}

	c.HTML(http.StatusOK, "layout", gin.H{
		"Requests":        requests,
		"ActivePage":      "worker-requests",
		"PageTitle":       "Worker Requests",
		"ContentTemplate": "worker_requests_content",
	})
}

func ApproveWorkerRequestHandler(c *gin.Context) {
	idParam := c.Param("id")

	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "DB bağlantı hatası: %v", err)
		return
	}
	defer dbConn.Close()

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.String(http.StatusBadRequest, "Geçersiz ID: %v", err)
		return
	}

	if err := worker_join_requests.ApproveWorkerJoinRequest(dbConn, id); err != nil {
		c.String(http.StatusInternalServerError, "Onaylama hatası: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/worker-requests")
}
