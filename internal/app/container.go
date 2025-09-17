package app

import (
	"umkm-api/config"
	"umkm-api/internal/umkm/handler"
	"umkm-api/internal/umkm/model"
	"umkm-api/internal/umkm/repository"
	"umkm-api/internal/umkm/service"
)

type Container struct {
	UmkmHandler *handler.UmkmHandler
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(&model.Umkm{})

	umkmRepo := repository.NewUmkmRepository(db)
	umkmService := service.NewUmkmService(umkmRepo)
	umkmHandler := handler.NewUmkmHandler(umkmService)

	return &Container{
		UmkmHandler: umkmHandler,
	}
}
