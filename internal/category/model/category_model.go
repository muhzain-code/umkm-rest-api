package model

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"size:50;not null" json:"name"`
	Photo *string `gorm:"default:null" json:"photo"`
	IsActive bool `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}