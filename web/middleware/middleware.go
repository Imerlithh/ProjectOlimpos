package middleware

import (
	"ProjectOlimpos/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RedirectToFirstRunIfNeeded() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !config.IsReady() && c.Request.URL.Path != "/first-run" {
			c.Redirect(http.StatusFound, "/first-run")
			c.Abort()
			return
		}
		c.Next()
	}
}
