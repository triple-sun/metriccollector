package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/triple-sun/metriccollector/internal/logger"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		logger.Log.
			Info().
			Str("method", r.Method).
			Str("url", r.URL.RequestURI()).
			Dur("ms", time.Since(time.Now())).
			Msg("incoming request")

		next.ServeHTTP(ww, r)

		logger.Log.
			Info().
			Int("status", ww.Status()).
			Int("size", ww.BytesWritten()).
			Msg("response")
	})
}
