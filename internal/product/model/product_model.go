package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	categoryModel "umkm-api/internal/category/model"
	umkmModel "umkm-api/internal/umkm/model"
)

type Product struct {
	ID     uuid.UUID      `gorm:"type:char(36);primaryKey;not null" json:"id"`
	UmkmID uuid.UUID      `gorm:"type:char(36);index;not null" json:"umkm_id"`
	Umkm   umkmModel.Umkm `gorm:"foreignKey:UmkmID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Umkm"`

	CategoryID int64                  `gorm:"index;not null" json:"category_id"`
	Category   categoryModel.Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Category"`

	Name        string         `gorm:"size:100;not null" json:"name"`
	Description *string        `gorm:"type:text" json:"description"`
	Price       float64        `gorm:"type:decimal(15,2);not null" json:"price"`
	Status      string         `gorm:"type:enum('available', 'pre_order', 'inactive');not null" json:"status"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	Photos       []ProductPhoto `gorm:"foreignKey:ProductID" json:"Photos"`
	Marketplaces []Marketplace  `gorm:"foreignKey:ProductID" json:"Marketplaces"`
}

type ProductPhoto struct {
	ProductID   uuid.UUID      `gorm:"type:char(36);not null" json:"product_id"`
	PhotoType   string         `gorm:"type:enum('utama', 'gallery');not null" json:"photo_type"`
	FilePath    string         `gorm:"size:255" json:"file_path"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	Description *string        `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Marketplace struct {
	ProductID       uuid.UUID      `gorm:"type:char(36);not null;index"`
	Name            string         `gorm:"size:50;not null"`
	Price           float64        `gorm:"type:decimal(15,2);not null"`
	MarketplaceLink string         `gorm:"size:255;not null"`
	IsActive        bool           `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
