package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/triple-sun/metriccollector/internal/handler"
	mware "github.com/triple-sun/metriccollector/internal/middleware"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func SetupRouter(storage *storage.MemStorage) *chi.Mux {
	r := chi.NewRouter()
	r.Use(mware.RequestLogger, mware.Gzip, middleware.Compress(5, "text/html", "application/json"))

	r.Get("/", handler.GetAllMetricsHandler(storage))

	r.Get("/value/{mtype}/{mname}", handler.GetMetricValueHandler(storage))
	r.Post("/update/{mtype}/{mname}/{mvalue}", handler.UpdateMetricHandler(storage))

	r.Post("/value", handler.GetMetricValueJSONHandler(storage))
	r.Post("/update", handler.UpdateMetricJSONHandler(storage))

	return r
}
