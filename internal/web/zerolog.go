package web

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/ftgogoV3/internal/logger"
	"github.com/rs/zerolog"
)

func ZeroLogger(zLogger zerolog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			ww := middleware.NewWrapResponseWriter(writer, request.ProtoMajor)

			start := time.Now()

			requestID := logger.GetRequestID(request.Context())
			correlationID := logger.GetCorrelationID(request.Context())
			causationID := logger.GetCausationID(request.Context())

			defer func() {
				var err error
				var logFn func() *zerolog.Event

				p := recover()

				switch {
				case p != nil:
					logFn = zLogger.Error().Stack
					// ensure the status code reflects this panic
					if ww.Status() < 500 {
						ww.WriteHeader(http.StatusInternalServerError)
					}
					err = errors.Errorf("%s", p)
				case ww.Status() < 400:
					logFn = zLogger.Info
				case ww.Status() < 500:
					logFn = zLogger.Warn
				default:
					logFn = zLogger.Error
				}
				log := logFn()
				if err != nil {
					log = log.Err(err)
				}
				log = log.Str("RemoteAddr", request.RemoteAddr).
					Int("ContextLength", ww.BytesWritten()).
					Dur("ResponseTime", time.Since(start))
				if requestID != "" {
					log = log.Str("RequestID", requestID).
						Str("CorrelationID", correlationID).
						Str("CausationID", causationID)
				}
				log.Msgf("[%d] %s %s", ww.Status(), request.Method, request.RequestURI)
			}()
			next.ServeHTTP(ww, request)
		})
	}
}
