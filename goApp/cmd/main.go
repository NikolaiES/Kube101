package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"sandvik.dev/goApp/internal/router"
)

func main() {
	godotenv.Load()
	print("Starting server")
	router.SetupRoutes()

	fmt.Print("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
