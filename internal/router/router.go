package router

import (
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/handler"
	"github.com/triple-sun/metriccollector/internal/middleware"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func SetupRouter(storage *storage.MemStorage) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.Use(middleware.RequestHandler(), middleware.ResponseHandler())

	r.GET("/", handler.GetAllMetricsHandler(storage))

	r.GET("/value/:mtype/:mname", handler.GetMetricValueHandler(storage))
	r.POST("/update/:mtype/:mname/:mvalue", handler.UpdateMetricHandler(storage))

	r.POST("/value", handler.GetMetricValueJSONHandler(storage))
	r.POST("/update", handler.UpdateMetricJSONHandler(storage))

	return r
}
