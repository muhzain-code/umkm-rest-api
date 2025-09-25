package repository

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"

	"gorm.io/gorm"
)

type ProductPromoRepository interface {
	Create(productPromo *model.ProductPromo) error
	FindAll(page, limit int, filter filter.ProductPromoFilter) ([]model.ProductPromo, int64, error)
	FindByID(id int64) (*model.ProductPromo, error)
	Update(productPromo *model.ProductPromo) error
	Delete(id int64) error
}

type productPromoRepository struct {
	db *gorm.DB
}

func NewProductPromoRepository(db *gorm.DB) ProductPromoRepository {
	return &productPromoRepository{db: db}
}

func (r *productPromoRepository) Create(productPromo *model.ProductPromo) error {
	return r.db.Create(productPromo).Error
}

func (r *productPromoRepository) FindAll(page, limit int, filter filter.ProductPromoFilter) ([]model.ProductPromo, int64, error) {
	var productPromos []model.ProductPromo
	var total int64

	if err := r.db.Model(&model.ProductPromo{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := r.db.Preload("Event").Preload("Product").Preload("Product.Umkm").Preload("Product.Category").Order("created_at desc").Limit(limit).Offset(offset).Find(&productPromos).Error
	return productPromos, total, err

}

func (r *productPromoRepository) FindByID(id int64) (*model.ProductPromo, error) {
	var productPromo model.ProductPromo
	err := r.db.First(&productPromo, id).Error
	if err != nil {
		return nil, err
	}
	return &productPromo, nil
}

func (r *productPromoRepository) Update(productPromo *model.ProductPromo) error {
	return r.db.Save(productPromo).Error
}

func (r *productPromoRepository) Delete(id int64) error {
	return r.db.Delete(&model.ProductPromo{}, id).Error
}
