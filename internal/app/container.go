package app

import (
	"umkm-api/config"
	categoryHandler "umkm-api/internal/category/handler"
	categoryModel "umkm-api/internal/category/model"
	categoryRepository "umkm-api/internal/category/repository"
	categoryService "umkm-api/internal/category/service"
	eventHandler "umkm-api/internal/event/handler"
	eventModel "umkm-api/internal/event/model"
	eventRepository "umkm-api/internal/event/repository"
	eventService "umkm-api/internal/event/service"
	productHandler "umkm-api/internal/product/handler"
	productModel "umkm-api/internal/product/model"
	productRepository "umkm-api/internal/product/repository"
	productService "umkm-api/internal/product/service"
	umkmHandler "umkm-api/internal/umkm/handler"
	umkmModel "umkm-api/internal/umkm/model"
	umkmRepository "umkm-api/internal/umkm/repository"
	umkmService "umkm-api/internal/umkm/service"
)

type Container struct {
	UmkmHandler     *umkmHandler.UmkmHandler
	CategoryHandler *categoryHandler.CategoryHandler
	ProductHandler  *productHandler.ProductHandler
	EventHandler    *eventHandler.EventHandler
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(&umkmModel.Umkm{}, &categoryModel.Category{}, &productModel.Product{}, 
		&productModel.ProductPhoto{}, &productModel.Marketplace{}, &eventModel.Event{}, &eventModel.EventUmkm{} )

	umkmRepo := umkmRepository.NewUmkmRepository(db)
	umkmService := umkmService.NewUmkmService(umkmRepo)
	umkmHandler := umkmHandler.NewUmkmHandler(umkmService)

	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepo)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)

	productRepo := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepo)
	productHandler := productHandler.NewProductHandler(productService)

	eventRepo := eventRepository.NewEventRepository(db)
	eventService := eventService.NewEventService(eventRepo)
	eventHandler := eventHandler.NewEventHandler(eventService)

	return &Container{
		UmkmHandler:     umkmHandler,
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
		EventHandler:    eventHandler,
	}
}
