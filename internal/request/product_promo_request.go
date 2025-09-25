package request

import (
	"github.com/google/uuid"
)

type CreateProductPromoRequest struct {
	EventID         uuid.UUID `json:"event_id" binding:"required,uuid"`
	ProductID       uuid.UUID `json:"product_id" binding:"required,uuid"`
	PromoPrice      float64   `json:"promo_price" binding:"required,gt=0"`
	DiscountAmount  *float64  `json:"discount_amount" binding:"omitempty,gt=0"`
	DiscountPercent *float64  `json:"discount_percent" binding:"omitempty,gt=0,lt=100"`
	IsActive        *bool     `json:"is_active" binding:"omitempty"`
	Terms           *string   `json:"terms" binding:"omitempty"`
}
type UpdateProductPromoRequest struct {
	EventID         uuid.UUID `json:"event_id" binding:"required,uuid"`
	ProductID       uuid.UUID `json:"product_id" binding:"required,uuid"`
	PromoPrice      float64   `json:"promo_price" binding:"required,gt=0"`
	DiscountAmount  *float64  `json:"discount_amount" binding:"omitempty,gt=0"`
	DiscountPercent *float64  `json:"discount_percent" binding:"omitempty,gt=0,lt=100"`
	IsActive        *bool     `json:"is_active" binding:"omitempty"`
	Terms           *string   `json:"terms" binding:"omitempty"`
}
