package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Umkm struct {
	ID           uuid.UUID      `gorm:"type:char(36);primaryKey" json:"id"`
	Name         string         `gorm:"size:50;not null" json:"name"`
	OwnerName    string         `gorm:"size:50;not null" json:"owner_name"`
	Nik          string         `gorm:"size:16;not null;uniqueIndex" json:"nik"`
	Gender       string         `gorm:"type:enum('l','p');not null" json:"gender"`
	Description  *string        `gorm:"type:text" json:"description"`
	PhotoProfile *string        `gorm:"size:255" json:"photo_profile"`
	Address      string         `gorm:"type:text;not null" json:"address"`
	Phone        string         `gorm:"size:20;not null;uniqueIndex" json:"phone"`
	Email        *string        `gorm:"size:50;uniqueIndex" json:"email"`
	WaLink       string         `gorm:"size:255; not null" json:"wa_link"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
