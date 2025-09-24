package app

import (
	"umkm-api/config"
	categoryHandler "umkm-api/internal/handler"
	categoryModel "umkm-api/internal/model"
	categoryRepository "umkm-api/internal/repository"
	categoryService "umkm-api/internal/service"
	eventHandler "umkm-api/internal/handler"
	eventModel "umkm-api/internal/model"
	eventRepository "umkm-api/internal/repository"
	eventService "umkm-api/internal/service"
	productHandler "umkm-api/internal/handler"
	productModel "umkm-api/internal/model"
	productRepository "umkm-api/internal/repository"
	productService "umkm-api/internal/service"
	umkmHandler "umkm-api/internal/handler"
	umkmModel "umkm-api/internal/model"
	umkmRepository "umkm-api/internal/repository"
	umkmService "umkm-api/internal/service"

	userHandler "umkm-api/internal/handler"
	userModel "umkm-api/internal/model"
	userRepository "umkm-api/internal/repository"
	userService "umkm-api/internal/service"

	"umkm-api/internal/auth"
)

type Container struct {
	UmkmHandler     *umkmHandler.UmkmHandler
	CategoryHandler *categoryHandler.CategoryHandler
	ProductHandler  *productHandler.ProductHandler
	EventHandler    *eventHandler.EventHandler
	UserHandler     *userHandler.UserHandler
	JWTService    auth.JWTService
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(&umkmModel.Umkm{}, &categoryModel.Category{}, &productModel.Product{},
		&productModel.ProductPhoto{}, &productModel.Marketplace{}, &eventModel.Event{}, &userModel.User{}, &eventModel.EventUmkm{})

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

	userRepo := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	jwtService := auth.NewJWTService()
	userHandler := userHandler.NewHandler(userService, jwtService)

	return &Container{
		UmkmHandler:     umkmHandler,
		CategoryHandler: categoryHandler,
		ProductHandler:  productHandler,
		EventHandler:    eventHandler,
		UserHandler:     userHandler,
		JWTService:    jwtService,
	}
}
