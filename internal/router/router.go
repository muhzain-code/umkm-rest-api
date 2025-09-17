package router

import (
	"github.com/gin-gonic/gin"
	"umkm-api/internal/umkm/handler"
	"umkm-api/internal/middleware"
)

func SetupRouter(
	umkmHandler *handler.UmkmHandler,
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

	return r
}
