package web

import (
	apiLogsGet "ProjectOlimpos/web/api_calls/api_logs/get"
	//apiLogsPost "ProjectOlimpos/web/api_calls/api_logs/post"
	serversGet "ProjectOlimpos/web/api_calls/servers/get"
	serversPost "ProjectOlimpos/web/api_calls/servers/post"
	workerRequestsGet "ProjectOlimpos/web/api_calls/worker_requests/get"
	//workerRequestsPost "ProjectOlimpos/web/api_calls/worker_requests/post"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Static files
	r.Static("/static", "./web/static")

	// Templates
	r.LoadHTMLGlob("web/templates/*.tmpl")

	// API Calls
	r.GET("/servers", serversGet.ListServersHandler)
	r.GET("/worker-requests", workerRequestsGet.ListWorkerRequestsHandler)
	r.GET("/api-logs", apiLogsGet.ListAPIRequestLogsHandler)
	r.GET("/servers/add", serversGet.ShowAddServerForm)
	r.POST("/servers/add", serversPost.AddServerHandler)
	return r
}
