package repository

import (
	"umkm-api/internal/model"

	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Create(app *model.Application) error
	FindAll(page, limit int, name string) ([]model.Application, int64, error)
	FindByID(id int64) (*model.Application, error)
	Update(app *model.Application) error
	Delete(id int64) error
}

type ApplicationRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicationRepository(db *gorm.DB) ApplicationRepository {
	return &ApplicationRepositoryImpl{db: db}
}

func (r *ApplicationRepositoryImpl) Create(app *model.Application) error {
	return r.db.Create(app).Error
}

func (r *ApplicationRepositoryImpl) FindAll(page, limit int, name string) ([]model.Application, int64, error) {
	var apps []model.Application
	var total int64

	query := r.db.Model(&model.Application{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&apps).Error; err != nil {
		return nil, 0, err
	}

	return apps, total, nil
}

func (r *ApplicationRepositoryImpl) FindByID(id int64) (*model.Application, error) {
	var app model.Application
	err := r.db.First(&app, id).Error
	if err != nil {
		return nil, err
	}

	return &app, nil
}
func (r *ApplicationRepositoryImpl) Update(app *model.Application) error {
	return r.db.Save(app).Error
}

func (r *ApplicationRepositoryImpl) Delete(id int64) error {
	return r.db.Delete(&model.Application{}, id).Error
}