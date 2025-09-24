package repository

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) error 
	FindAll(page, limit int, filter filter.CategoryFilter) ([]model.Category, int64, error)
	FindByID(id int64) (*model.Category, error)
	Update(category *model.Category) error
	Delete(id int64) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(category *model.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) FindAll(page, limit int, ft filter.CategoryFilter) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	db := r.db.Model(&model.Category{})
	db = filter.ApplyCategoryFilter(db, ft)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Order("created_at desc").Limit(limit).Offset(offset).Find(&categories).Error

	return categories, total, err
}

func (r *categoryRepository) FindByID(id int64) (*model.Category, error) {
	var category model.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) Update(category *model.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) Delete(id int64) error {
	return r.db.Delete(&model.Category{}, id).Error
}