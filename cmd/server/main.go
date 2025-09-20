package main

import (
	"log"
	"os"
	"umkm-api/internal/app"
	"umkm-api/internal/router"
)

func main() {
	container := app.BuildContainer()

	r := router.SetupRouter(
		container.UmkmHandler,
		container.CategoryHandler,
		container.EventHandler,
	)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("APP_HOST")
	if host == "" {
		host = "localhost"
	}

	log.Printf("Starting server at http://%s:%s", host, port)

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
