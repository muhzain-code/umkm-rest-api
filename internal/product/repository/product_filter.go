package repository

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductFilter struct {
	UmkmID     uuid.UUID
	CategoryID int64
	Name     string
	IsActive *bool
}

func ApplyProductFilter(db *gorm.DB, filter ProductFilter) *gorm.DB {
	if strings.TrimSpace(filter.Name) != "" {
		db = db.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *filter.IsActive)
	}

	if filter.UmkmID != uuid.Nil {
		db = db.Where("umkm_id = ?", filter.UmkmID)
	}

	if filter.CategoryID != 0 {
		db = db.Where("category_id = ?", filter.CategoryID)
	}

	return db
}
