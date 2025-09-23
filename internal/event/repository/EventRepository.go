package repository

import (
	"umkm-api/internal/event/model"

	"gorm.io/gorm"
)

type EventRepository interface{
	FindAll(page, limit int)([]model.Event,int64 ,error)
	FindByID(id int)(*model.Event, error)
	Create(event *model.Event) error
	Update(umkm *model.Event) error
	Delete(id int) error
}

type eventRepository struct{
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db: db}
}

func (m *eventRepository) FindAll(page, limit int) ([]model.Event, int64, error) {
	var Marketplace []model.Event
	var total int64

	db := m.db.Model(&model.Event{})
	if err := db.Count(&total).Error; err != nil {
		return  nil, 0, err
	} 

	Offset := (page - 1) * limit
	err := db.Order("created_at desc").Limit(limit).Offset(Offset).Find(&Marketplace).Error

	return Marketplace, total, err
}

func (i *eventRepository) FindByID(id int)(*model.Event, error)  {
	var event model.Event
	err := i.db.First(&event, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (C *eventRepository) Create(event *model.Event) error {
	return C.db.Create(event).Error
}

func (u *eventRepository) Update(event *model.Event) error {
	return u.db.Save(event).Error
}

func (v *eventRepository) Delete(id int) error {
	return v.db.Delete(&model.Event{}, "id = ?",id).Error
}