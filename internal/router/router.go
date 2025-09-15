package router

import (
	"github.com/gin-gonic/gin"
	"umkm-api/internal/handler"
)

func SetupRouter(
	umkmHandler *handler.UmkmHandler,
) *gin.Engine {
	r := gin.Default()

	umkm := r.Group("/umkms") 
	{
		umkm.GET("/", umkmHandler.GetAllUmkm)
		umkm.GET("/:id", umkmHandler.GetUmkmByID)
		umkm.POST("/", umkmHandler.CreateUmkm)
		umkm.PUT("/:id", umkmHandler.UpdateUmkm)
		umkm.DELETE("/:id", umkmHandler.DeleteUmkm)
	}

	return r
}