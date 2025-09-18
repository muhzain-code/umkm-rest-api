package repository

import (
	"gorm.io/gorm"
	"strings"
)

type CategoryFilter struct {
	Name string
	IsActive *bool
}

func ApplyCategoryFilter(db *gorm.DB, filter CategoryFilter) *gorm.DB {
	if strings.TrimSpace(filter.Name) != "" {
		db = db.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *&filter.IsActive)
	}
	return db
}