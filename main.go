/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-15 21:47:55
 * @LastEditTime: 2019-08-15 22:17:04
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"net/http"
	"github.com/oxygen-org/server/api/apiv1"
	"github.com/oxygen-org/server/api/apiv2"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	 "github.com/oxygen-org/server/admin"
)


func setupRouter() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	admin.InitAdmin(e)
	// Ping test
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	api1_0 := e.Group("/api/v1.0")
	api1_0.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "xtalpi" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))
	api2_0 := e.Group("/api/v2.0")

	{
		api1_0.GET("/version", apiv1.Version)

		api1_0.GET("/archive", apiv1.Archive)
		api1_0.POST("/archive", apiv1.Archive)
		api1_0.DELETE("/archive", apiv1.Archive)

		api1_0.GET("/archive/check", apiv1.ArchiveCheck)
		api1_0.POST("/archive/check", apiv1.ArchiveCheck)

		api1_0.GET("/archive/:archive_id/detail", apiv1.ArchiveDetail)
		api1_0.PATCH("/archive/:archive_id/detail", apiv1.ArchiveDetail)

		api1_0.GET("/datasets", apiv1.Datasets)
		api1_0.POST("/datasets", apiv1.Datasets)
		api1_0.DELETE("/datasets", apiv1.Datasets)

		api1_0.GET("/datasets/checking", apiv1.DatasetsCheck)

		api1_0.GET("/datasets/:data_id/detail", apiv1.DatasetsCheck)
		api1_0.PATCH("/datasets/:data_id/detail", apiv1.DatasetsCheck)

		api1_0.GET("/images", apiv1.Images)
		api1_0.POST("/images", apiv1.Images)

		api1_0.POST("/image/deploy", apiv1.ImagesDeploy)

		api1_0.GET("/image/deploy/confirmation", apiv1.ImageConfirmation)

		api1_0.Any("/jobs", apiv1.Jobs)

		api1_0.GET("/job/:job_id", apiv1.JobsDetail)

		api1_0.POST("/job/runner", apiv1.JobsRun)

		api1_0.POST("/job/killer", apiv1.JobsKill)

		api1_0.Any("/job/state", apiv1.JobState)

		api1_0.GET("/job/:job_id/log", apiv1.JobLog)

		api1_0.GET("/job/rpc_service", apiv1.JobRPCService)

		api1_0.DELETE("/job/log", apiv1.JobLogClear)

		api1_0.GET("/:job_id/events", apiv1.JobEvents)
		api1_0.POST("/:job_id/events", apiv1.JobEvents)

		api1_0.GET("/resources", apiv1.Resources)

		api1_0.GET("/monitor", apiv1.ResourceMonitor)

		api1_0.GET("/version", apiv1.Version)
		api1_0.GET("/version", apiv1.Version)
		api1_0.GET("/version", apiv1.Version)
		api1_0.GET("/version", apiv1.Version)
		api1_0.GET("/version", apiv1.Version)
	}
	{
		api2_0.GET("/version", apiv2.Version)
		api2_0.GET("/resource", apiv2.StatResource)
		api2_0.GET("/resource_cluster", apiv2.StatResourceCluster)
		api2_0.GET("/job", apiv2.StatJob)
		api2_0.GET("/job_distribute", apiv2.StateJobDistribute)
		api2_0.GET("/dataset", apiv2.StateDataset)
	}
	return e
}
