package router

import (
	"github.com/gin-gonic/gin"
	umkmHandler "umkm-api/internal/umkm/handler"
	categoryHandler "umkm-api/internal/category/handler"
	productHandler "umkm-api/internal/product/handler"
	eventHandler "umkm-api/internal/event/handler"
	"umkm-api/internal/middleware"
)

func SetupRouter(
	umkmHandler *umkmHandler.UmkmHandler,
	categoryHandler *categoryHandler.CategoryHandler,
	productHandler *productHandler.ProductHandler,
	eventHandler *eventHandler.EventHandler,
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

	umkm := r.Group("/umkms")
	{
		umkm.GET("", umkmHandler.GetAllUmkm)
		umkm.GET("/:id", umkmHandler.GetUmkmByID)
		umkm.POST("", umkmHandler.CreateUmkm)
		umkm.PUT("/:id", umkmHandler.UpdateUmkm)
		umkm.DELETE("/:id", umkmHandler.DeleteUmkm)
	}

	category := r.Group("/categories")
	{
		category.GET("", categoryHandler.GetAllCategory)
		category.GET("/:id", categoryHandler.GetCategoryByID)
		category.POST("", categoryHandler.CreateCategory)
		category.PUT("/:id", categoryHandler.UpdateCategory)
		category.DELETE("/:id", categoryHandler.DeleteCategory)
	}

	product := r.Group("/products")
	{
		product.GET("", productHandler.GetAllProducts)
		product.GET("/:id", productHandler.GetProductByID)
		product.POST("", productHandler.CreateProduct)
		product.PUT("/:id", productHandler.UpdateProduct)
		product.DELETE("/:id", productHandler.DeleteProduct)
	}

	event := r.Group("/events")
	{
		event.GET("", eventHandler.GetAllEvent)
		event.GET("/:id", eventHandler.GetEventByID)
		event.POST("", eventHandler.CreateEvent)
		event.PUT("/:id", eventHandler.UpdateEvent)
		event.DELETE("/:id", eventHandler.DeleteEvent)
	}
	return r
}