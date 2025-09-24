package filter

import (
	"strings"

	"gorm.io/gorm"
)

type UmkmFilter struct {
	Name   string
	IsActive *bool
}

func ApplyUmkmFilter(db *gorm.DB, filter UmkmFilter) *gorm.DB {
	if strings.TrimSpace(filter.Name) != "" {
		db = db.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *filter.IsActive)
	}

	return db
}
