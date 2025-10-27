package router

import (
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/handler"
	"github.com/triple-sun/metriccollector/internal/middleware"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func SetupRoutes(storage *storage.MemStorage) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.Use(middleware.ResponseLogger())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.RequestLogger())

	r.GET("/", handler.GetAllMetricsHandler(storage))
	r.POST("/value", handler.MetricGetValueHandler(storage))
	r.POST("/update", handler.MetricUpdateHandler(storage))

	return r
}
