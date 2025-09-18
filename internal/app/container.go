package app

import (
	"umkm-api/config"
	umkmHandler "umkm-api/internal/umkm/handler"
	umkmModel "umkm-api/internal/umkm/model"
	umkmRepository "umkm-api/internal/umkm/repository"
	umkmService "umkm-api/internal/umkm/service"
	categoryHandler "umkm-api/internal/category/handler"
	categoryModel "umkm-api/internal/category/model"
	categoryRepository "umkm-api/internal/category/repository"
	categoryService "umkm-api/internal/category/service"
)

type Container struct {
	UmkmHandler *umkmHandler.UmkmHandler
	CategoryHandler *categoryHandler.CategoryHandler
}

func BuildContainer() *Container {
	db := config.ConnectDB()

	db.AutoMigrate(&umkmModel.Umkm{}, &categoryModel.Category{})

	umkmRepo := umkmRepository.NewUmkmRepository(db)
	umkmService := umkmService.NewUmkmService(umkmRepo)
	umkmHandler := umkmHandler.NewUmkmHandler(umkmService)

	categoryRepo := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepo)
	categoryHandler := categoryHandler.NewCategoryHandler(categoryService)

	return &Container{
		UmkmHandler: umkmHandler,
		CategoryHandler: categoryHandler,
	}
}
