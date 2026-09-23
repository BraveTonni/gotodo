package core_http_middleware

import (
	"context"
	"net/http"
	"time"

	core_logger "github.com/BraveTonni/gotodo/internal/core/logger"
	core_http_response "github.com/BraveTonni/gotodo/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func RequestID() Middleware {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				requestID := r.Header.Get(requestIDHeader)

				if requestID == "" {
					requestID = uuid.NewString()
				}

				r.Header.Set(requestIDHeader, requestID)
				w.Header().Set(requestIDHeader, requestID)

				next.ServeHTTP(w, r)
			})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "unexpected panic got during handle HTTP request")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Metrics() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)

			start_time := time.Now()
			log.Debug(
				">>> incoming HTTP REQUEST",
				zap.Time("time", start_time.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				">>> done HTTP REQUEST",
				zap.Int("status_code", rw.GetStatusCodeOrPanic()),
				zap.Duration("latency", time.Since(start_time)),
			)
		})
	}
}
