package server

import (
	"gohook/internal/config"
	"gohook/internal/server/routes"
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", routes.Health)
	mux.HandleFunc("POST /{id}/{token}", routes.Webhook)

	err := http.ListenAndServe(config.Get().Address, mux)
	if err != nil {
		log.Fatal("Failed to start server:", err.Error())
	}
}
