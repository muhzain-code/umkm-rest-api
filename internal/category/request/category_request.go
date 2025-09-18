package request

import (
	"mime/multipart"
	"umkm-api/pkg/utils"
)

type CreateCategoryRequest struct {
	Name      string                `form:"name" binding:"required,max=50"`
	Photo     *multipart.FileHeader `form:"photo"`
	PhotoPath *string               `json:"-"`
}

type UpdateCategoryRequest struct {
	Name      string                `form:"name" binding:"required,max=50"`
	Photo     *multipart.FileHeader `form:"photo"`
	PhotoPath *string               `json:"-"`
	IsActive  *bool                 `form:"is_active"`
}

func (r *CreateCategoryRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.Photo, false)
}

func (r *UpdateCategoryRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.Photo, false)
}
