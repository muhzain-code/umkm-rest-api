package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string         `gorm:"size:100;index;not null;comment:Nama event" json:"name"`
	Description string         `gorm:"type:text;comment:Deskripsi event" json:"description"`
	Photo       *string         `gorm:"size:255;comment:Foto event" json:"photo"`
	StartDate   time.Time      `gorm:"type:date;not null;comment:Tanggal mulai" json:"start_date"`
	EndDate     time.Time      `gorm:"type:date;not null;comment:Tanggal selesai" json:"end_date"`
	IsActive    bool           `gorm:"default:true;index;comment:Status aktif" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}