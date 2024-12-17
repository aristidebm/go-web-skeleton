package main

import (
	"log"
	"log/slog"
	"net/http"
)

var ListenAddr = ":8001"

func main() {
	handler := RegisterURLs()
	slog.Info("The server has started and is running on", "address", ListenAddr)
	log.Fatal(http.ListenAndServe(ListenAddr, handler))
}
