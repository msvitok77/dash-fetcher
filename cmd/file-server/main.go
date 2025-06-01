package main

import (
	"log/slog"
	"net/http"
)

func main() {
	fileHandler := http.FileServer(http.Dir("./mpds"))

	go func() {
		slog.Info("🚀 https handler started")
		err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", fileHandler)
		if err != nil {
			slog.Error("https handler failed", "err", err)
		}
	}()

	slog.Info("🚀 http handler started")
	err := http.ListenAndServe(":8080", fileHandler)
	if err != nil {
		slog.Info("http handler failed", "err", err)
	}
}
