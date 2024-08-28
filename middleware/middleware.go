package middleware

import (
	"log/slog"
	"net/http"
)

func New(h http.Handler) func(*slog.Logger) http.Handler {
	return func(l *slog.Logger) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Info("Middleware",
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
