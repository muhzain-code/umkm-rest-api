package repository

import (
	"fmt"
	"os"
	"path/filepath"
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) error
	FindAll(page, limit int, filter filter.ProductFilter) ([]model.Product, int64, error)
	FindByID(id uuid.UUID) (*model.Product, error)
	Update(product *model.Product, photos []*model.ProductPhoto, marketplaces []*model.Marketplace) error
	Delete(id uuid.UUID) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var umkms *model.Umkm
		if err := tx.First(&umkms, "id = ?", product.UmkmID).Error; err != nil {
			fmt.Println("error find umkm:", err)
			return err
		}

		var category *model.Category
		if err := tx.First(&category, "id = ?", product.CategoryID).Error; err != nil {
			fmt.Println("error find category:", err)
			return err
		}

		if err := tx.Create(product).Error; err != nil {
			return err
		}

		if err := tx.Preload("Umkm").
			Preload("Category").
			Preload("Photos").
			Preload("Marketplaces").
			First(product, "id = ?", product.ID).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *productRepository) FindAll(page, limit int, ft filter.ProductFilter) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	db := r.db.Model(&model.Product{})
	db = filter.ApplyProductFilter(db, ft)

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.
		Preload("Umkm").
		Preload("Category").
		Preload("Photos").
		Preload("Marketplaces").
		Order("created_at desc").
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, total, err
}

func (r *productRepository) FindByID(id uuid.UUID) (*model.Product, error) {
	var product model.Product
	if err := r.db.
		Preload("Umkm").
		Preload("Category").
		Preload("Photos").
		Preload("Marketplaces").
		First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *model.Product, photos []*model.ProductPhoto, marketplaces []*model.Marketplace) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(product).Error; err != nil {
			return err
		}

		if err := tx.Where("product_id = ?", product.ID).Delete(&model.ProductPhoto{}).Error; err != nil {
			return err
		}
		for _, p := range photos {
			p.ProductID = product.ID
		}
		if len(photos) > 0 {
			if err := tx.Create(&photos).Error; err != nil {
				return err
			}
		}

		if err := tx.Where("product_id = ?", product.ID).Delete(&model.Marketplace{}).Error; err != nil {
			return err
		}
		for _, m := range marketplaces {
			m.ProductID = product.ID
		}
		if len(marketplaces) > 0 {
			if err := tx.Create(&marketplaces).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *productRepository) Delete(id uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		var photos []model.ProductPhoto
		if err := tx.Where("product_id = ?", id).Find(&photos).Error; err != nil {
			return err
		}

		for _, p := range photos {
			if p.FilePath != "" {
				fullPath := filepath.Join("uploads", p.FilePath)
				if err := os.Remove(fullPath); err != nil && !os.IsNotExist(err) {
					fmt.Printf("failed to delete photo file %s: %v\n", fullPath, err)
				}
			}
		}

		if err := tx.Where("product_id = ?", id).Delete(&model.ProductPhoto{}).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id = ?", id).Delete(&model.Marketplace{}).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.Product{}, "id = ?", id).Error; err != nil {
			return err
		}

		return nil
	})
}
