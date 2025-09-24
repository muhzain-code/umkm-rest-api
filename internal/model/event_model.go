package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID          uuid.UUID      `gorm:"type:char(36);primaryKey;not null" json:"id"`
	Name        string         `gorm:"size:100;index;not null;comment:Nama event" json:"name"`
	Description string         `gorm:"type:text;comment:Deskripsi event" json:"description"`
	Photo       *string        `gorm:"size:255;comment:Foto event" json:"photo"`
	StartDate   time.Time      `gorm:"type:date;not null;comment:Tanggal mulai" json:"start_date"`
	EndDate     time.Time      `gorm:"type:date;not null;comment:Tanggal selesai" json:"end_date"`
	IsActive    bool           `gorm:"default:true;index;comment:Status aktif" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	EventUmkms []EventUmkm `gorm:"foreignKey:EventID" json:"event_umkms"`
}

type EventUmkmResponse struct {
	UmkmID   uuid.UUID `json:"umkm_id"`
	IsActive bool      `json:"is_active"`
}

type EventResponse struct {
	ID          uuid.UUID           `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Photo       *string             `json:"photo"`
	StartDate   time.Time           `json:"start_date"`
	EndDate     time.Time           `json:"end_date"`
	IsActive    bool                `json:"is_active"`
	EventUmkms  []EventUmkmResponse `json:"event_umkms"`
}
type EventUmkm struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	UmkmID    uuid.UUID      `gorm:"type:char(36);not null;uniqueIndex:uq_event_umkm_umkm_id_event_id;comment:Relasi ke UMKM" json:"umkm_id"`
	EventID   uuid.UUID      `gorm:"type:char(36);not null;uniqueIndex:uq_event_umkm_umkm_id_event_id;comment:Relasi ke Event" json:"event_id"`
	IsActive  bool           `gorm:"default:true;index;comment:Status aktif" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Umkm  Umkm  `gorm:"foreignKey:UmkmID;constraint:OnDelete:CASCADE;" json:"umkm"`
	Event Event `gorm:"foreignKey:EventID;constraint:OnDelete:CASCADE;" json:"event"`
}

func (Event) TableName() string {
	return "events"
}
func (EventUmkm) TableName() string {
	return "event_umkm"
}
