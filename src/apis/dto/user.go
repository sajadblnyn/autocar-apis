package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,len=11,iran-mobile-validator"`
}
