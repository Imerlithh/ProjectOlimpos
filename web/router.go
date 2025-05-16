package web

import (
	"ProjectOlimpos/web/api_calls"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Static files
	r.Static("/static", "./web/static")

	// Templates
	r.LoadHTMLGlob("web/templates/*.tmpl")

	// API Calls
	r.GET("/servers", api_calls.ListServersHandler)
	r.GET("/worker-requests", api_calls.ListWorkerRequestsHandler)
	r.GET("/api-logs", api_calls.ListAPIRequestLogsHandler)
	r.GET("/servers/add", api_calls.ShowAddServerForm)
	r.POST("/servers/add", api_calls.AddServerHandler)
	return r
}
