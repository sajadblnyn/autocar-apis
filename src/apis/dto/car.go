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

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,max=30,min=3"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearboxId" binding:"required"`
}
type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty" binding:"max=30,min=3"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}
type CarModelResponse struct {
	Id             int                     `json:"id"`
	Name           string                  `json:"name"`
	CompanyId      int                     `json:"companyId"`
	CarTypeId      int                     `json:"carTypeId"`
	GearboxId      int                     `json:"gearboxId"`
	Company        CompanyResponse         `json:"company"`
	Gearbox        GearboxResponse         `json:"gearbox"`
	CarType        CarTypeResponse         `json:"carType"`
	CarModelColors []CarModelColorResponse `json:"carModelColors,omitempty"`
}

type CreateCarModelColorRequest struct {
	ColorId    int `json:"colorId" binding:"required"`
	CarModelId int `json:"carModelId" binding:"required"`
}
type UpdateCarModelColorRequest struct {
	ColorId    int `json:"colorId,omitempty"`
	CarModelId int `json:"carModelId,omitempty"`
}

type CarModelColorResponse struct {
	ColorId    int           `json:"colorId"`
	Color      ColorResponse `json:"color"`
	CarModelId int           `json:"carModelId"`
}
