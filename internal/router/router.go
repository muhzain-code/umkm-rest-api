package router

import (
	"github.com/gin-gonic/gin"
	categoryHandler "umkm-api/internal/handler"
	eventHandler "umkm-api/internal/handler"
	"umkm-api/internal/middleware"
	productHandler "umkm-api/internal/handler"
	umkmHandler "umkm-api/internal/handler"
	"umkm-api/internal/auth"
	userHandler "umkm-api/internal/handler"
	logHistoryHandler "umkm-api/internal/handler"
)

func SetupRouter(
	umkmHandler *umkmHandler.UmkmHandler,
	categoryHandler *categoryHandler.CategoryHandler,
	productHandler *productHandler.ProductHandler,
	eventHandler *eventHandler.EventHandler,
	jwtService auth.JWTService,
	userHandler *userHandler.UserHandler,
	logHistoryHandler *logHistoryHandler.LogHistoryHandler,
) *gin.Engine {
	r := gin.New()

	r.Static("/uploads", "./uploads")

	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.Recovery())
	r.Use(middleware.RequestID())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	r.POST("/api/register", userHandler.Register)
	r.POST("/api/login", userHandler.Login)
	r.POST("/api/log", logHistoryHandler.Create)

	api := r.Group("/api")
	api.Use(auth.JWTAuthMiddleware(jwtService))
	{
		umkm := api.Group("/umkms")
		{
			umkm.GET("", umkmHandler.GetAllUmkm)
			umkm.GET("/:id", umkmHandler.GetUmkmByID)
			umkm.POST("", umkmHandler.CreateUmkm)
			umkm.PUT("/:id", umkmHandler.UpdateUmkm)
			umkm.DELETE("/:id", umkmHandler.DeleteUmkm)
		}

		category := api.Group("/categories")
		{
			category.GET("", categoryHandler.GetAllCategory)
			category.GET("/:id", categoryHandler.GetCategoryByID)
			category.POST("", categoryHandler.CreateCategory)
			category.PUT("/:id", categoryHandler.UpdateCategory)
			category.DELETE("/:id", categoryHandler.DeleteCategory)
		}

		product := api.Group("/products")
		{
			product.GET("", productHandler.GetAllProducts)
			product.GET("/:id", productHandler.GetProductByID)
			product.POST("", productHandler.CreateProduct)
			product.PUT("/:id", productHandler.UpdateProduct)
			product.DELETE("/:id", productHandler.DeleteProduct)
		}

		event := api.Group("/events")
		{
			event.GET("", eventHandler.GetAllEvent)
			event.GET("/:id", eventHandler.GetEventByID)
			event.POST("", eventHandler.CreateEvent)
			event.PUT("/:id", eventHandler.UpdateEvent)
			event.DELETE("/:id", eventHandler.DeleteEvent)
		}
	}

	return r
}
