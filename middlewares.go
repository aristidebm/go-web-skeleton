package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
)

type ResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (lw *ResponseWriter) WriteHeader(statusCode int) {
	lw.StatusCode = statusCode
	lw.ResponseWriter.WriteHeader(statusCode)
}

func LoggingResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lw := &ResponseWriter{ResponseWriter: w, StatusCode: 200}
		next.ServeHTTP(lw, r)
		slog.Info("", "method", r.Method, "path", r.URL.Path, "code", lw.StatusCode, "status", http.StatusText(lw.StatusCode))
	})
}

func PanicRecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		defer func() {
			if err := recover(); err != nil {
				slog.Error("Something went wrong", "error", err)
				slog.Debug(stackTrace(err.(error)))
				// TODO: Show the stack trace if DEBUG=true
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
	})
}

func stackTrace(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%s\n%s", err, debug.Stack())
}
