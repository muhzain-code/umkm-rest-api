package model

import (
	"gorm.io/gorm"
	"time"
)

type ActivityLog struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	SessionID *string        `gorm:"size:64;index" json:"session_id"`
	ProductID string        `gorm:"type:char(36);index" json:"product_id"`
	LogType   string         `gorm:"type:enum('view','buy');index" json:"log_type"`
	IPAddress *string        `gorm:"size:45" json:"ip_address"`
	UserAgent *string        `gorm:"size:255" json:"user_agent"`
	Referrer  *string        `gorm:"size:255" json:"referrer"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
