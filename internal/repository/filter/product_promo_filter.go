package filter

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductPromoFilter struct {
	EventID   uuid.UUID
	ProductID uuid.UUID
	IsActive  *bool
}

func ApplyProductPromoFilter(db *gorm.DB, filter ProductPromoFilter) *gorm.DB {
	if filter.EventID != uuid.Nil {
		db = db.Where("event_id = ?", filter.EventID)
	}
	if filter.ProductID != uuid.Nil {
		db = db.Where("product_id = ?", filter.ProductID)
	}
	if filter.IsActive != nil {
		db = db.Where("is_active = ?", *filter.IsActive)
	}
	return db
}
