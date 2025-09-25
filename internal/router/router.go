package router

import (
	"github.com/gin-gonic/gin"
	"umkm-api/internal/auth"
	"umkm-api/internal/handler"
	"umkm-api/internal/middleware"
)

func SetupRouter(
	umkmHandler *handler.UmkmHandler,
	categoryHandler *handler.CategoryHandler,
	productHandler *handler.ProductHandler,
	eventHandler *handler.EventHandler,
	jwtService auth.JWTService,
	userHandler *handler.UserHandler,
	logHistoryHandler *handler.LogHistoryHandler,
	AppHandler *handler.ApplicationHandler,
	ProductPromoHandler *handler.ProductPromoHandler,
) *gin.Engine {
	r := gin.New()

	// Serve static files
	r.Static("/uploads", "./uploads")

	// Global middlewares
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())
	r.Use(middleware.Recovery())
	r.Use(middleware.RequestID())

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server is running!"})
	})

	// Public routes
	r.POST("/api/register", userHandler.Register)
	r.POST("/api/login", userHandler.Login)
	r.POST("/api/log", logHistoryHandler.Create)

	// Protected routes
	api := r.Group("/api")
	{
		// UMKM routes
		umkm := api.Group("/umkms")
		{
			umkm.GET("", umkmHandler.GetAllUmkm)
			umkm.GET("/:id", umkmHandler.GetUmkmByID)
			umkm.POST("", umkmHandler.CreateUmkm).Use(auth.JWTAuthMiddleware(jwtService))
			umkm.PUT("/:id", umkmHandler.UpdateUmkm).Use(auth.JWTAuthMiddleware(jwtService))
			umkm.DELETE("/:id", umkmHandler.DeleteUmkm).Use(auth.JWTAuthMiddleware(jwtService))
		}

		// Category routes
		category := api.Group("/categories")
		{
			category.GET("", categoryHandler.GetAllCategory)
			category.GET("/:id", categoryHandler.GetCategoryByID)
			category.POST("", categoryHandler.CreateCategory).Use(auth.JWTAuthMiddleware(jwtService))
			category.PUT("/:id", categoryHandler.UpdateCategory).Use(auth.JWTAuthMiddleware(jwtService))
			category.DELETE("/:id", categoryHandler.DeleteCategory).Use(auth.JWTAuthMiddleware(jwtService))
		}

		// Product routes
		product := api.Group("/products")
		{
			product.GET("", productHandler.GetAllProducts)
			product.GET("/:id", productHandler.GetProductByID)
			product.POST("", productHandler.CreateProduct).Use(auth.JWTAuthMiddleware(jwtService))
			product.PUT("/:id", productHandler.UpdateProduct).Use(auth.JWTAuthMiddleware(jwtService))
			product.DELETE("/:id", productHandler.DeleteProduct).Use(auth.JWTAuthMiddleware(jwtService))
		}

		// Event routes
		event := api.Group("/events")
		{
			event.GET("", eventHandler.GetAllEvent)
			event.GET("/:id", eventHandler.GetEventByID)
			event.POST("", eventHandler.CreateEvent).Use(auth.JWTAuthMiddleware(jwtService))
			event.PUT("/:id", eventHandler.UpdateEvent).Use(auth.JWTAuthMiddleware(jwtService))
			event.DELETE("/:id", eventHandler.DeleteEvent).Use(auth.JWTAuthMiddleware(jwtService))
		}

		// Applications
		app := api.Group("/applications")
		{
			app.GET("", AppHandler.GetAllApplication)
			app.GET("/:id", AppHandler.GetApplicationByID)
			app.POST("", AppHandler.CreateApplication).Use(auth.JWTAuthMiddleware(jwtService))
			app.PUT("/:id", AppHandler.UpdateApplication).Use(auth.JWTAuthMiddleware(jwtService))
			app.DELETE("/:id", AppHandler.DeleteApplication).Use(auth.JWTAuthMiddleware(jwtService))
		}

		// Product promos
		productPromo := api.Group("/product-promos")
		{
			productPromo.GET("", ProductPromoHandler.GetAllProductPromo)
			productPromo.GET("/:id", ProductPromoHandler.GetProductPromoByID)
			productPromo.POST("", ProductPromoHandler.CreateProductPromo).Use(auth.JWTAuthMiddleware(jwtService))
			productPromo.PUT("/:id", ProductPromoHandler.UpdateProductPromo).Use(auth.JWTAuthMiddleware(jwtService))
			productPromo.DELETE("/:id", ProductPromoHandler.DeleteProductPromo).Use(auth.JWTAuthMiddleware(jwtService))
		}
	}

	return r
}
