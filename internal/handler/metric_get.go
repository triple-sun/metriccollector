package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/triple-sun/metriccollector/internal/requests"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/storage"
)

func GetAllMetricsHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result string

		fmt.Println("Got all metrics")

		metrics := storage.GetMetrics()

		for name, metric := range *metrics {
			result += fmt.Sprintf("<b>%s</b>: %s<br/>", name, metric.GetValue())
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(200)
		if _, err := w.Write(fmt.Appendf(nil, `<html><body>%s</body></html>`, result)); err != nil {
			w.WriteHeader(500)
		}
	}

}

func GetMetricValueJSONHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.GetMetricValueRequest{}

		if err := render.Bind(r, req); err != nil {
			_ = render.Render(w, r, responses.NewErrorResponse(400, err))
			return
		}

		if metric, err := storage.FindOne(req); err != nil {
			_ = render.Render(w, r, responses.NewErrorResponse(404, err))
			return
		} else {
			render.JSON(w, r, metric)
		}
	}
}

func GetMetricValueHandler(storage *storage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := &requests.GetMetricValueRequest{MType: chi.URLParam(r, "mtype"), ID: chi.URLParam(r, "mname")}

		if err := req.Validate(); err != nil {
			_ = render.Render(w, r, responses.NewErrorResponse(400, err))
			return
		}

		if metric, err := storage.FindOne(req); err != nil {
			_ = render.Render(w, r, responses.NewErrorResponse(404, err))
		} else {
			render.JSON(w, r, metric)
		}
	}
}
