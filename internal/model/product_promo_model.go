package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ProductPromo struct {
	ID              int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	EventID         uuid.UUID      `gorm:"type:char(36);not null;index;comment:Relasi ke Event" json:"event_id"`
	Event           Event          `gorm:"foreignKey:EventID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"event"`
	ProductID       uuid.UUID      `gorm:"type:char(36);not null;index;comment:Relasi ke Product" json:"product_id"`
	Product         Product        `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	PromoPrice      float64        `gorm:"type:decimal(15,2);not null;comment:Harga promo" json:"promo_price"`
	DiscountAmount  *float64       `gorm:"type:decimal(15,2);comment:Jumlah diskon" json:"discount_amount"`
	DiscountPercent *float64       `gorm:"type:decimal(5,2);comment:Persentase diskon" json:"discount_percent"`
	IsActive        bool           `gorm:"default:true;index;comment:Status aktif" json:"is_active"`
	Terms           *string        `gorm:"type:text;comment:Syarat dan ketentuan promo" json:"terms"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
