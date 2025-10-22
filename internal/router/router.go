package router

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/handler"
	"github.com/triple-sun/metriccollector/internal/middleware"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func SetupRoutes(storage *storage.MemStorage) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(requestid.New())
	r.Use(middleware.ErrorHandler())

	r.GET("/", handler.GetAllMetricsHandler(storage))
	r.GET("/value/:mtype/:mname", handler.MetricGetValueHandler(storage))

	r.POST("/update/:mtype/:mname/:mvalue", handler.MetricUpdateHandler(storage))

	return r
}
