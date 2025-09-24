package repository

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UmkmRepository interface {
	Create(umkm *model.Umkm) error
	FindAll(page, limit int, filter filter.UmkmFilter) ([]model.Umkm, int64, error)
	FindByID(id uuid.UUID) (*model.Umkm, error)
	Update(umkm *model.Umkm) error
	Delete(id uuid.UUID) error
	
}

type umkmRepository struct {
	db *gorm.DB
}

func NewUmkmRepository(db *gorm.DB) UmkmRepository {
	return &umkmRepository{db: db}
}

func (r *umkmRepository) Create(umkm *model.Umkm) error {
	return r.db.Create(umkm).Error
}

func (r *umkmRepository) FindAll(page, limit int, ft filter.UmkmFilter) ([]model.Umkm, int64, error) {
	var umkms []model.Umkm
	var total int64

	db := r.db.Model(&model.Umkm{})
	db = filter.ApplyUmkmFilter(db, ft)
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Order("created_at desc").Limit(limit).Offset(offset).Find(&umkms).Error

	return umkms, total, err
}

func (r *umkmRepository) FindByID(id uuid.UUID) (*model.Umkm, error) {
	var umkm model.Umkm
	err := r.db.First(&umkm, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &umkm, nil
}

func (r *umkmRepository) Update(umkm *model.Umkm) error {
	return r.db.Save(umkm).Error
}

func (r *umkmRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&model.Umkm{}, "id = ?", id).Error
}
