package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	"github.com/triple-sun/metriccollector/internal/handler"
	"github.com/triple-sun/metriccollector/internal/middleware"
	"github.com/triple-sun/metriccollector/internal/storage"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func SetupRoutes(storage *storage.MemStorage) *gin.Engine {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.Use(middleware.ResponseLogger())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.RequestLogger())

	r.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithDecompressFn(gzip.DefaultDecompressHandle), gzip.WithCustomShouldCompressFn(utils.ShouldServerCompress)))

	r.GET("/", handler.GetAllMetricsHandler(storage))

	r.GET("/value/:mtype/:mname", handler.MetricGetValueHandler(storage))
	r.POST("/update/:mtype/:mname/:mvalue", handler.MetricUpdateHandler(storage))

	r.POST("/value", handler.MetricGetValueJSONHandler(storage))
	r.POST("/update", handler.MetricUpdateJSONHandler(storage))

	return r
}
