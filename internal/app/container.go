package app

import (
	"umkm-api/config"
	categoryHandler "umkm-api/internal/handler"
	eventHandler "umkm-api/internal/handler"
	productHandler "umkm-api/internal/handler"
	umkmHandler "umkm-api/internal/handler"
	categoryModel "umkm-api/internal/model"
	eventModel "umkm-api/internal/model"
	productModel "umkm-api/internal/model"
	umkmModel "umkm-api/internal/model"
	categoryRepository "umkm-api/internal/repository"
	eventRepository "umkm-api/internal/repository"
	productRepository "umkm-api/internal/repository"
	umkmRepository "umkm-api/internal/repository"
	categoryService "umkm-api/internal/service"
	eventService "umkm-api/internal/service"
	productService "umkm-api/internal/service"
	umkmService "umkm-api/internal/service"

	userHandler "umkm-api/internal/handler"
	userModel "umkm-api/internal/model"
	userRepository "umkm-api/internal/repository"
	userService "umkm-api/internal/service"

	logHistoryHandler "umkm-api/internal/handler"
	logHistoryModel "umkm-api/internal/model"
	logHistoryRepository "umkm-api/internal/repository"
	logHistoryService "umkm-api/internal/service"

	activityLogModel "umkm-api/internal/model"
	activityLogRepository "umkm-api/internal/repository"
	activityLogService "umkm-api/internal/service"

	"umkm-api/internal/auth"
)

type Container struct {
	UmkmHandler        *umkmHandler.UmkmHandler
	CategoryHandler    *categoryHandler.CategoryHandler
	ProductHandler     *productHandler.ProductHandler
	EventHandler       *eventHandler.EventHandler
	UserHandler        *userHandler.UserHandler
	JWTService         auth.JWTService
	LogHistoryHandler  *logHistoryHandler.LogHistoryHandler
	ActivityLogService activityLogService.ActivityLogService
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(&umkmModel.Umkm{}, &categoryModel.Category{}, &productModel.Product{},
		&productModel.ProductPhoto{}, &productModel.Marketplace{}, &eventModel.Event{}, &userModel.User{}, &eventModel.EventUmkm{}, &activityLogModel.ActivityLog{}, &logHistoryModel.LogHistory{})

	activityLogRepository := activityLogRepository.NewActivityLogRepository(db)
	activityLog := activityLogService.NewActivityLogService(activityLogRepository)

	umkmRepo := umkmRepository.NewUmkmRepository(db)
	umkmService := umkmService.NewUmkmService(umkmRepo)
	umkmHandler := umkmHandler.NewUmkmHandler(umkmService)

	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepo)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)

	productRepo := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepo)
	productHandler := productHandler.NewProductHandler(productService, activityLog)

	eventRepo := eventRepository.NewEventRepository(db)
	eventService := eventService.NewEventService(eventRepo)
	eventHandler := eventHandler.NewEventHandler(eventService)

	userRepo := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	jwtService := auth.NewJWTService()
	userHandler := userHandler.NewHandler(userService, jwtService)

	logHistoryRepo := logHistoryRepository.NewLogHistoryRepository(db)
	logHistoryService := logHistoryService.NewLogHistoryService(logHistoryRepo)
	logHistoryHandler := logHistoryHandler.NewLogHistoryHandler(logHistoryService, activityLog)

	return &Container{
		UmkmHandler:        umkmHandler,
		CategoryHandler:    categoryHandler,
		ProductHandler:     productHandler,
		EventHandler:       eventHandler,
		UserHandler:        userHandler,
		JWTService:         jwtService,
		LogHistoryHandler:  logHistoryHandler,
		ActivityLogService: activityLog,
	}
}
