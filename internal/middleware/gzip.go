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

			defer func() {
				err := cr.Close()
				if err != nil {
					logger.Log.Error().Err(err)
				}
			}()

			r.Body = cr
		}

		/**
		acceptsGzip := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")

		if (isJSON || isHTML) && acceptsGzip {
			logger.Log.Info().Msg("Compressing response body...")

			cw := utils.NewCompressWriter(w)

			ow = cw

			defer func() {
				err := cw.Close()
				if err != nil {
					logger.Log.Error().Err(err)
				}
				logger.Log.Info().Msg("Response body compressed!")
			}()
		} */

		next.ServeHTTP(ow, r)
	})
}
