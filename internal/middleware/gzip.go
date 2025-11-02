package middleware

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"

	"github.com/triple-sun/metriccollector/internal/logger"
	"github.com/triple-sun/metriccollector/internal/responses"
	"github.com/triple-sun/metriccollector/internal/utils"
)

func Gzip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ow := w

		isJSON := strings.Contains(r.Header.Get("Content-Type"), "application/json")
		isHTML := strings.Contains(r.Header.Get("Content-Type"), "text/html")

		sendsGzip := strings.Contains(r.Header.Get("Content-Encoding"), "gzip")

		if (isHTML || isJSON) && sendsGzip {
			logger.Log.Info().Msg("Decompressing request body...")

			cr, err := utils.NewCompressReader(r.Body)

			if err != nil {
				render.JSON(w, r, responses.NewErrorResponse(500, err))
				return
			}

			defer cr.Close()

			r.Body = cr
		}

		next.ServeHTTP(ow, r)
	})
}
