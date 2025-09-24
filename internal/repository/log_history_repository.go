package repository

import (
	"umkm-api/internal/model"
	"gorm.io/gorm"
)

type LogHistoryRepository interface {
	Create(log *model.LogHistory) error
}

type logHistoryRepository struct {
	db *gorm.DB
}

func NewLogHistoryRepository(db *gorm.DB) LogHistoryRepository {
	return &logHistoryRepository{db: db}
}

func (r *logHistoryRepository) Create(log *model.LogHistory) error {
	return r.db.Create(log).Error
}