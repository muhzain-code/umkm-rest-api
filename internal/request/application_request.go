package request

import "umkm-api/pkg/utils"

type CreateApplicationRequest struct {
	Name string `json:"name" binding:"required,max=100"`
}

type UpdateApplicationRequest struct {
	Name string `json:"name" binding:"required,max=100"`
}

func (r *CreateApplicationRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return nil
}

func (r *UpdateApplicationRequest) Validate() error {
	if err := utils.Validate.Struct(r); err != nil {
		return err
	}
	return nil
}