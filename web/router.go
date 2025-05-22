package web

import (
	apiLogsGet "ProjectOlimpos/web/api_calls/api_logs/get"
	apiLogsPost "ProjectOlimpos/web/api_calls/api_logs/post"
	serversGet "ProjectOlimpos/web/api_calls/servers/get"
	serversPost "ProjectOlimpos/web/api_calls/servers/post"
	workerRequestsGet "ProjectOlimpos/web/api_calls/worker_requests/get"
	"ProjectOlimpos/web/middleware"

	//workerRequestsPost "ProjectOlimpos/web/api_calls/worker_join_requests/post"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Static files
	r.Static("/static", "./web/static")

	// Templates
	r.LoadHTMLGlob("web/templates/*.tmpl")
	r.Use(middleware.RedirectToFirstRunIfNeeded())

	r.GET("/first-run", ShowFirstRunForm)
	r.POST("/first-run", HandleFirstRunSubmit)
	r.GET("/", HomePage)
	// API Calls

	//Servers
	r.GET("/servers", serversGet.ListServersHandler)
	r.GET("/servers/add", serversGet.ShowAddServerForm)
	r.POST("/servers/add", serversPost.AddServerHandler)

	//Worker Join Requests
	r.GET("/worker-requests", workerRequestsGet.ListWorkerJoinRequestsHandler)

	//API logs
	r.GET("/api-logs", apiLogsGet.ListAPIRequestLogsHandler)
	r.POST("/api-logs/delete/:id", apiLogsPost.DeleteAPIRequestLogHandler)

	return r
}
