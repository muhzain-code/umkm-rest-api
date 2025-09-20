package router

import (
	"github.com/gin-gonic/gin"
	categoryHandler "umkm-api/internal/category/handler"
	eventHandler "umkm-api/internal/event/handler"
	"umkm-api/internal/middleware"
	umkmHandler "umkm-api/internal/umkm/handler"
)

func SetupRouter(
	umkmHandler *umkmHandler.UmkmHandler,
	categoryHandler *categoryHandler.CategoryHandler,
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