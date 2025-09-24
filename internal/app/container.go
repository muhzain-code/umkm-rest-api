package app

import (
	"umkm-api/config"
	"umkm-api/internal/auth"

	"umkm-api/internal/handler"
	"umkm-api/internal/model"
	"umkm-api/internal/repository"
	"umkm-api/internal/service"
)

type Container struct {
	UmkmHandler        *handler.UmkmHandler
	CategoryHandler    *handler.CategoryHandler
	ProductHandler     *handler.ProductHandler
	EventHandler       *handler.EventHandler
	UserHandler        *handler.UserHandler
	JWTService         auth.JWTService
	LogHistoryHandler  *handler.LogHistoryHandler
	ActivityLogService service.ActivityLogService
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(
		&model.Umkm{},
		&model.Category{},
		&model.Product{},
		&model.ProductPhoto{},
		&model.Marketplace{},
		&model.Event{},
		&model.User{},
		&model.EventUmkm{},
		&model.ActivityLog{},
		&model.LogHistory{},
	)

	// ActivityLog
	activityLogRepo := repository.NewActivityLogRepository(db)
	activityLogSvc := service.NewActivityLogService(activityLogRepo)

	// Umkm
	umkmRepo := repository.NewUmkmRepository(db)
	umkmSvc := service.NewUmkmService(umkmRepo)
	umkmH := handler.NewUmkmHandler(umkmSvc)

	// Category
	categoryRepo := repository.NewCategoryRepository(db)
	categorySvc := service.NewCategoryService(categoryRepo)
	categoryH := handler.NewCategoryHandler(categorySvc)

	// Product
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(productRepo)
	productH := handler.NewProductHandler(productSvc, activityLogSvc)

	// Event
	eventRepo := repository.NewEventRepository(db)
	eventSvc := service.NewEventService(eventRepo)
	eventH := handler.NewEventHandler(eventSvc)

	// User
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(userRepo)
	jwtSvc := auth.NewJWTService()
	userH := handler.NewHandler(userSvc, jwtSvc)

	// LogHistory
	logHistoryRepo := repository.NewLogHistoryRepository(db)
	logHistorySvc := service.NewLogHistoryService(logHistoryRepo)
	logHistoryH := handler.NewLogHistoryHandler(logHistorySvc, activityLogSvc)

	return &Container{
		UmkmHandler:        umkmH,
		CategoryHandler:    categoryH,
		ProductHandler:     productH,
		EventHandler:       eventH,
		UserHandler:        userH,
		JWTService:         jwtSvc,
		LogHistoryHandler:  logHistoryH,
		ActivityLogService: activityLogSvc,
	}
}
