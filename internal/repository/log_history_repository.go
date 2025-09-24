package repository

import (
	"umkm-api/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LogHistoryRepository interface {
	Create(log *model.LogHistory) error
	FindUmkmByProduct(id uuid.UUID) (uuid.UUID, error)
	CountByResi(resi string) (int64, error)
}

type logHistoryRepository struct {
	db *gorm.DB
}

func NewLogHistoryRepository(db *gorm.DB) LogHistoryRepository {
	return &logHistoryRepository{db: db}
}

func (r *logHistoryRepository) FindUmkmByProduct(id uuid.UUID) (uuid.UUID, error) {
	var product model.Product
	err := r.db.Select("umkm_id").First(&product, "id = ?", id).Error
	if err != nil {
		return uuid.Nil, err 
	}
	return product.UmkmID, nil
}

func (r *logHistoryRepository) Create(log *model.LogHistory) error {
	return r.db.Create(log).Error
}

func (r *logHistoryRepository) CountByResi(resi string) (int64, error) {
	var count int64
	err := r.db.Model(&model.LogHistory{}).Where("resi = ?", resi).Count(&count).Error
	return count, err
}
