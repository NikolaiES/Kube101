package router

import (
	"net/http"
	"sandvik.dev/goApp/internal/handlers"
)

func SetupRoutes() {

	http.HandleFunc("GET /secret", handlers.HandleSecret)
	http.HandleFunc("GET /data", handlers.HandleLorem)
	http.HandleFunc("GET /healthz", handlers.ReportHealth)
	http.HandleFunc("POST /healthz", handlers.SetHealth)
	http.HandleFunc("GET /spacebarheating", handlers.SpaceBarHeating)
}
