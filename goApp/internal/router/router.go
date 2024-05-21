package router

import (
	"net/http"
	"sandvik.dev/goApp/internal/handlers"
)

func SetupRoutes() {

	http.HandleFunc("GET /secret", handlers.HandleSecret)
}
