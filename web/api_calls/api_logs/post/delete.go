package post

import (
	"net/http"
	"strconv"

	"ProjectOlimpos/db"
	gormlogs "ProjectOlimpos/db/operations/gorm/api_logs"
	"github.com/gin-gonic/gin"
)

func DeleteAPIRequestLogHandler(c *gin.Context) {
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Geçersiz ID")
		return
	}
	id := uint(id64)

	if err := gormlogs.DeleteAPIRequestLogByID(db.DB, id); err != nil {
		c.String(http.StatusInternalServerError, "Silinemedi: %v", err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/api-logs")
}
