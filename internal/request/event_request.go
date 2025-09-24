package request

import (
	"mime/multipart"
	"umkm-api/pkg/utils"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Name        string                `form:"name" binding:"required,max=100"`
	Description string                `form:"description" binding:"required"`
	Photo       *multipart.FileHeader `form:"photo"`
	PhotoPath   *string               `json:"-"`
	StartDate   string                `form:"start_date" binding:"required"`
	EndDate     string                `form:"end_date" binding:"required"`
	UmkmID      string                `form:"umkm_id" binding:"required"`
}

type UpdateEventRequest struct {
	Name        string                `form:"name" binding:"required,max=100"`
	Description string                `form:"description" binding:"required"`
	Photo       *multipart.FileHeader `form:"photo"`
	PhotoPath   *string               `json:"-"`
	StartDate   string                `form:"start_date" binding:"required"`
	EndDate     string                `form:"end_date" binding:"required"`
	IsActive    *bool                 `form:"is_active"`

	UmkmIDs []uuid.UUID `form:"umkm_ids[]" json:"umkm_ids"`
}

func (r *UpdateEventRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.Photo, false)
}

func (r *CreateEventRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.Photo, true)
}
