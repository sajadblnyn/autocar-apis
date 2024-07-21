package dto

type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,max=30,alpha,min=3"`
}
type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,max=30,alpha,min=3"`
}
type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,max=30,alpha,min=3"`
}
type UpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,max=30,alpha,min=3"`
}
type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type CreateCompanyRequest struct {
	Name      string `json:"name" binding:"required,max=30,alpha,min=3"`
	CountryId int    `json:"countryId" binding:"required"`
}
type UpdateCompanyRequest struct {
	Name      string `json:"name,omitempty" binding:"max=30,alpha,min=3"`
	CountryId int    `json:"countryId,omitempty"`
}

type CompanyResponse struct {
	Name      string             `json:"name"`
	CountryId int                `json:"countryId"`
	Country   GetCountryResponse `json:"country,omitempty"`
}
