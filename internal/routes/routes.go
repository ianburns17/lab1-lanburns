package routes

import (
	"net/http"

	"github.com/ianburns17/lab2-Ianburns/internal/handlers"
	"github.com/ianburns17/lab2-Ianburns/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)
	mux.Handle("/hello", middleware.Logging(http.HandlerFunc(handlers.Hello)))
	return mux
}
