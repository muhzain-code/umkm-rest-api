package main

import (
	"log"
	"os"
	"umkm-api/internal/app"
	"umkm-api/internal/router"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"umkm-api/internal/request"
)

func main() {
	container := app.BuildContainer()

	r := router.SetupRouter(
		container.UmkmHandler,
		container.CategoryHandler,
		container.ProductHandler,
		container.EventHandler,
		container.JWTService,
		container.UserHandler,
		container.LogHistoryHandler,
		container.ApplicationHandler,
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("photos", request.ValidatePhotos)
		v.RegisterValidation("marketplaces", request.ValidateMarketplaces)
	}

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
