package repository

import (
	"umkm-api/internal/model"
	"umkm-api/internal/repository/filter"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	FindAll(page, limit int, filter filter.EventFilter) ([]model.Event, int64, error)
	FindByID(id int) (*model.Event, error)
	Create(event *model.Event, umkmIDs []uuid.UUID) error
	Update(umkm *model.Event) error
	Delete(id int) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (m *eventRepository) FindAll(page, limit int, ft filter.EventFilter) ([]model.Event, int64, error) {
	var events []model.Event
	var total int64

	db := m.db.Model(&model.Event{})
	db = filter.ApplyEventFilter(db, ft)

	// hitung total
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := db.Preload("EventUmkms", func(db *gorm.DB) *gorm.DB {
		return db.Select("umkm_id", "is_active", "event_id")
	}).
		Order("events.created_at desc").
		Limit(limit).Offset(offset).Find(&events).Error

	// err := db.
	// 	Preload("EventUmkms", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("id", "event_id", "umkm_id") // ambil kolom penting saja
	// 	}).
	// Order("events.created_at desc").
	// Limit(limit).Offset(offset).
	// 	Find(&events).Error

	return events, total, err
}

func (i *eventRepository) FindByID(id int) (*model.Event, error) {
	var event model.Event
	err := i.db.Preload("EventUmkms").Preload("EventUmkms.Umkm").First(&event, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *eventRepository) Create(event *model.Event, umkmIDs []uuid.UUID) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Insert ke tabel events
		if err := tx.Create(event).Error; err != nil {
			return err
		}

		// 2. Insert ke tabel event_umkm
		for _, umkmID := range umkmIDs {
			eventUmkm := model.EventUmkm{
				EventID:  event.ID,
				UmkmID:   umkmID,
				IsActive: true,
			}
			if err := tx.Create(&eventUmkm).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (u *eventRepository) Update(event *model.Event) error {
	return u.db.Save(event).Error
}

func (v *eventRepository) Delete(id int) error {
	return v.db.Delete(&model.Event{}, "id = ?", id).Error
}
