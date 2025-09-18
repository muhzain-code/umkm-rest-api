package request

import (
	"mime/multipart"
	"umkm-api/pkg/utils"
 
)

type CreateUmkmRequest struct {
	Name             string                `form:"name" binding:"required,max=50"`
	OwnerName        string                `form:"owner_name" binding:"required,max=50"`
	Nik              string                `form:"nik" binding:"required,len=16"`
	Gender           string                `form:"gender" binding:"required,oneof=l p"`
	Description      *string               `form:"description"`
	PhotoProfile     *multipart.FileHeader `form:"photo_profile"`
	PhotoProfilePath *string               `json:"-"`
	Address          string                `form:"address" binding:"required"`
	Phone            string                `form:"phone" binding:"required,max=20"`
	Email            *string               `form:"email"`
	WaLink           string                `form:"wa_link" binding:"required"`
}

type UpdateUmkmRequest struct {
	Name             string                `form:"name" binding:"omitempty,max=50"`
	OwnerName        string                `form:"owner_name" binding:"omitempty,max=50"`
	Nik              string                `form:"nik" binding:"required,len=16"`
	Gender           string                `form:"gender" binding:"omitempty,oneof=l p"`
	Description      *string               `form:"description"`
	PhotoProfile     *multipart.FileHeader `form:"photo_profile"`
	PhotoProfilePath *string               `json:"-"`
	Address          string                `form:"address" binding:"omitempty"`
	Phone            string                `form:"phone" binding:"omitempty,max=20"`
	Email            *string               `form:"email"`
	IsActive         *bool                 `form:"is_active"`
	WaLink           string                `form:"wa_link" binding:"required"`
}

func (r *CreateUmkmRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.PhotoProfile, true)
}

func (r *UpdateUmkmRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return utils.ValidatePhoto(r.PhotoProfile, false)
}
