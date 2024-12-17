package main

import (
	"net/http"
)

func RegisterURLs() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /hello/{$}", Make(Hello))
	mux.Handle("GET /todos/{$}", Make(ListTodos))
	mux.Handle("GET /login/{$}", Make(Login))
	return PanicRecoveryMiddleware(LoggingResponseMiddleware(mux))
}
