package model

import (
	"time"

	"gorm.io/gorm"
)

type Application struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	Name      string `gorm:"size:100;unique;not null;comment:Nama aplikasi"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Aplication []LogHistory `gorm:"foreignKey:ApplicationID" json:"Aplication"`
}
