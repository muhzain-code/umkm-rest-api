package app

import (
	"umkm-api/config"
	"umkm-api/internal/handler"
	"umkm-api/internal/repository"
	"umkm-api/internal/service"
	"umkm-api/internal/model"
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