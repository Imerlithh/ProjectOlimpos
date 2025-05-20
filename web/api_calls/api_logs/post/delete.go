package api_logs

import (
	"net/http"
	"strconv"

	"ProjectOlimpos/config"
	"ProjectOlimpos/db"
	"ProjectOlimpos/db/operations/tables/api_logs"
	"github.com/gin-gonic/gin"
)

func DeleteAPIRequestLogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Geçersiz ID")
		return
	}

	dbConn, err := db.Connect(config.Load())
	if err != nil {
		c.String(http.StatusInternalServerError, "Veritabanı bağlantı hatası")
		return
	}
	defer dbConn.Close()

	err = api_logs.DeleteAPIRequestLogByID(dbConn, id)
	if err != nil {
		c.String(http.StatusInternalServerError, "Kayıt silinemedi: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/api-logs")
}
