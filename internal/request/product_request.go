package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type PhotoRequest struct {
	PhotoType   string  `json:"photo_type" binding:"required,oneof=utama gallery"`
	FilePath    string  `json:"file_path,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	Description *string `json:"description,omitempty"`
}

type MarketplaceRequest struct {
	Name            string  `json:"name" binding:"required,max=50"`
	Price           float64 `json:"price" binding:"required,gte=0"`
	MarketplaceLink string  `json:"marketplace_link" binding:"required,max=255"`
	IsActive        *bool   `json:"is_active,omitempty"`
}

type CreateProductRequest struct {
	UmkmID       string               `form:"umkm_id" binding:"required"`
	CategoryID   int64                `form:"category_id" binding:"required"`
	Name         string               `form:"name" binding:"required"`
	Description  *string              `form:"description"`
	Price        float64              `form:"price" binding:"required"`
	Status       string               `form:"status" binding:"required,oneof=available pre_order inactive"`
	Photos       []PhotoRequest       `json:"photos"`       
	Marketplaces []MarketplaceRequest `json:"marketplaces"` 
}

type UpdateProductRequest struct {
	UmkmID       uuid.UUID            `form:"umkm_id" binding:"required"`
	CategoryID   int64                `form:"category_id" binding:"required"`
	Name         string               `form:"name" binding:"required"`
	Description  *string              `form:"description"`
	Price        float64              `form:"price" binding:"required"`
	Status       string               `form:"status" binding:"required,oneof=available pre_order inactive"`
	Photos       []PhotoRequest       `json:"photos"`
	Marketplaces []MarketplaceRequest `json:"marketplaces"`
}

func ValidatePhotos(fl validator.FieldLevel) bool {
	photos, ok := fl.Field().Interface().([]PhotoRequest)
	if !ok {
		return false
	}

	countUtama := 0
	for _, p := range photos {
		if p.PhotoType == "utama" {
			countUtama++
		}
	}

	return countUtama == 1
}

func ValidateMarketplaces(fl validator.FieldLevel) bool {
	marketplaces, ok := fl.Field().Interface().([]MarketplaceRequest)
	if !ok {
		return false
	}

	for _, m := range marketplaces {
		if m.Name != "" || m.Price > 0 || m.MarketplaceLink != "" {
			if m.Name == "" || m.Price <= 0 || m.MarketplaceLink == "" {
				return false
			}
		}
	}

	return true
}
