package request

type CreateUmkmRequest struct {
	Name         string  `json:"name" binding:"required,max=50"`
	OwnerName    string  `json:"owner_name" binding:"required,max=50"`
	Nik          string  `json:"nik" binding:"required,len=16"`
	Gender       string  `json:"gender" binding:"required,oneof=l p"`
	Description  *string `json:"description"`
	PhotoProfile *string `json:"photo_profile"`
	Address      string  `json:"address" binding:"required"`
	Phone        string  `json:"phone" binding:"required,max=20"`
	Email        *string `json:"email"`
	WaLink       string  `json:"wa_link" binding:"required"`
}

type UpdateUmkmRequest struct {
	Name         string  `json:"name" binding:"omitempty,max=50"`
	OwnerName    string  `json:"owner_name" binding:"omitempty,max=50"`
	Nik          string  `json:"nik" binding:"required,len=16"`
	Gender       string  `json:"gender" binding:"omitempty,oneof=l p"`
	Description  *string `json:"description"`
	PhotoProfile *string `json:"photo_profile"`
	Address      string  `json:"address" binding:"omitempty"`
	Phone        string  `json:"phone" binding:"omitempty,max=20"`
	Email        *string `json:"email"`
	IsActive     *bool   `json:"is_active"`
	WaLink       string  `json:"wa_link" binding:"required"`
}
