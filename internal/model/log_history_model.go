package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type LogHistory struct {
	ApplicationID int         `gorm:"index;not null" json:"application_id"`
	Application   Application `gorm:"foreignKey:ApplicationID;references:ID" json:"application"`
	UmkmID        uuid.UUID   `gorm:"type:char(36);index;not null" json:"umkm_id"`
	Umkm          Umkm        `gorm:"foreignKey:UmkmID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Umkm"`

	Resi         string         `gorm:"char:17;not null" json:"resi"`
	BuyerName    string         `gorm:"size:100;not null" json:"buyer_name"`
	BuyerPhone   string         `gorm:"size:20;not null" json:"buyer_phone"`
	BuyerAddress string         `gorm:"type:text;not null" json:"buyer_address"`
	IpAddress    string         `gorm:"size:45;not null" json:"ip_address"`
	UserAgent    string         `gorm:"type:text; not null" json:"user_agent"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
