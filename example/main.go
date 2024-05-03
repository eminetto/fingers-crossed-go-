package main

import (
	"log/slog"
	"net/http"
	"os"

	middleware "github.com/eminetto/fingers-crossed-go"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/info", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Info("info inside the handler")
		w.Write([]byte("Hello World with info"))
	})
	r.Get("/error", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Error("error inside the handler")
		w.Write([]byte("Hello World with error"))
	})
	r.Get("/debug", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Debug("debug inside the handler")
		w.Write([]byte("Hello World with debug"))
	})
	r.Get("/warn", func(w http.ResponseWriter, r *http.Request) {
		logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
		logger.Warn("warn inside the handler")
		w.Write([]byte("Hello World with warn"))
	})
	var logs []middleware.LogEntry
	fg := middleware.FingersCrossed(slog.LevelInfo, slog.LevelError, logs, r)
	http.ListenAndServe(":3000", fg)
}