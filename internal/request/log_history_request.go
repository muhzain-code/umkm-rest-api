package request

import "github.com/google/uuid"

type LogHistoryRequest struct {
	ApplicationID int `json:"application_id" binding:"required"`
	ProductID   uuid.UUID `json:"product_id" binding:"required"`
	BuyerName     string `json:"buyer_name" binding:"omitempty,max=100"`
	BuyerPhone    string `json:"buyer_phone" binding:"omitempty,max=20"`
	BuyerAddress  string `json:"buyer_address" binding:"omitempty,max=255"`
	IPAddress     string `json:"ip_address" binding:"omitempty,ipv4|ipv6"`
	UserAgent     string `json:"user_agent" binding:"omitempty,max=255"`
	Resi          string `json:"resi" binding:"omitempty,max=17"`
}
