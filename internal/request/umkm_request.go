package request

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"slices"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CreateUmkmRequest struct {
	Name             string                `form:"name" binding:"required,max=50"`
	OwnerName        string                `form:"owner_name" binding:"required,max=50"`
	Nik              string                `form:"nik" binding:"required,len=16"`
	Gender           string                `form:"gender" binding:"required,oneof=l p"`
	Description      *string               `form:"description"`
	PhotoProfile     *multipart.FileHeader `form:"photo_profile"` // hanya untuk binding
	PhotoProfilePath *string               `json:"-"`             // disiapkan handler, dipakai service
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

var validate = validator.New()

func (r *CreateUmkmRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return validatePhoto(r.PhotoProfile, true)
}

func (r *UpdateUmkmRequest) Validate() error {
	if err := validate.Struct(r); err != nil {
		return err
	}
	return validatePhoto(r.PhotoProfile, false)
}

func validatePhoto(photo *multipart.FileHeader, required bool) error {
	if photo == nil {
		if required {
			return errors.New("photo profile required")
		}
		return nil
	}

	ext := strings.ToLower(filepath.Ext(photo.Filename))
	allowed := []string{".jpg", ".png", ".jpeg"}
	if !slices.Contains(allowed, ext) {
		return errors.New("invalid photo format")
	}
	return nil
}
