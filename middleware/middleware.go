package middleware

import (
	"log/slog"
	"net/http"
)

func New(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Middleware",
				slog.String("Method", r.Method),
				slog.String("Host", r.Host),
				slog.String("Path", r.URL.Path),
				slog.String("Query", r.URL.RawQuery),
				slog.String("IP", r.RemoteAddr),
			)

			h.ServeHTTP(w, r)
		})
	}
}
