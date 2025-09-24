package filter

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventFilter struct {
	IsActive  *bool
	Search    string
	StartDate *time.Time
	EndDate   *time.Time
	UmkmID    *uuid.UUID
}

func ApplyEventFilter(db *gorm.DB, filter EventFilter) *gorm.DB {
	if filter.IsActive != nil {
		db = db.Where("events.is_active = ?", *filter.IsActive)
	}

	if filter.Search != "" {
		db = db.Where("events.name LIKE ?", "%"+filter.Search+"%")
	}

	if filter.StartDate != nil {
		db = db.Where("events.start_date >= ?", *filter.StartDate)
	}

	if filter.EndDate != nil {
		db = db.Where("events.end_date <= ?", *filter.EndDate)
	}

	if filter.UmkmID != nil {
		db = db.Joins("INNER JOIN event_umkm eu ON eu.event_id = events.id AND eu.umkm_id = ?", *filter.UmkmID)
	}

	return db
}
