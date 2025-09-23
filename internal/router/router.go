package router

import (
	"github.com/gin-gonic/gin"
	categoryHandler "umkm-api/internal/category/handler"
	eventHandler "umkm-api/internal/event/handler"
	"umkm-api/internal/middleware"
	productHandler "umkm-api/internal/product/handler"
	umkmHandler "umkm-api/internal/umkm/handler"
	"umkm-api/internal/user/auth"
	userHandler "umkm-api/internal/user/handler"
)

func SetupRouter(
	umkmHandler *umkmHandler.UmkmHandler,
	categoryHandler *categoryHandler.CategoryHandler,
	productHandler *productHandler.ProductHandler,
	eventHandler *eventHandler.EventHandler,
	jwtService auth.JWTService,
	userHandler *userHandler.UserHandler,
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

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

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
