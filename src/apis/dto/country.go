package dto

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,max=20,min=3"`
}

type GetCountryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
