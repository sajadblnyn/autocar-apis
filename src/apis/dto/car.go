package dto

import "time"

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
	Id                 int                           `json:"id"`
	Name               string                        `json:"name"`
	CompanyId          int                           `json:"companyId"`
	CarTypeId          int                           `json:"carTypeId"`
	GearboxId          int                           `json:"gearboxId"`
	Company            CompanyResponse               `json:"company"`
	Gearbox            GearboxResponse               `json:"gearbox"`
	CarType            CarTypeResponse               `json:"carType"`
	CarModelColors     []CarModelColorResponse       `json:"carModelColors,omitempty"`
	CarModelYears      []CarModelPersianYearResponse `json:"carModelYears,omitempty"`
	CarModelImages     []CarModelImageResponse       `json:"carModelImages,omitempty"`
	CarModelProperties []CarModelPropertyResponse    `json:"carModelProperties,omitempty"`
	CarModelComments   []CarModelCommentResponse     `json:"carModelComments,omitempty"`
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

type CreateCarModelPersianYearRequest struct {
	PersianYearId int `json:"persianYearId" binding:"required"`
	CarModelId    int `json:"carModelId" binding:"required"`
}
type UpdateCarModelPersianYearRequest struct {
	PersianYearId int `json:"persianYearId,omitempty"`
	CarModelId    int `json:"carModelId,omitempty"`
}

type CarModelPersianYearResponse struct {
	Id                     int                            `json:"id"`
	PersianYearId          int                            `json:"persianYearId"`
	PersianYear            PersianYearWithoutDateResponse `json:"persianYear"`
	CarModelId             int                            `json:"carModelId"`
	CarModelPriceHistories []CarModelPriceHistoryResponse `json:"carModelPriceHistories,omitempty"`
}

type CreateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	Price          float64   `json:"price" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
}
type UpdateCarModelPriceHistoryRequest struct {
	Price   float64   `json:"price,omitempty"`
	PriceAt time.Time `json:"priceAt,omitempty"`
}

type CarModelPriceHistoryResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId"`
	Price          float64   `json:"price"`
	PriceAt        time.Time `json:"priceAt"`
}

type CreateCarModelImageRequest struct {
	ImageId     int  `json:"imageId" binding:"required"`
	CarModelId  int  `json:"carModelId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}
type UpdateCarModelImageRequest struct {
	IsMainImage bool `json:"isMainImage"`
}

type CarModelImageResponse struct {
	Id          int             `json:"id"`
	ImageId     int             `json:"imageId"`
	CarModelId  int             `json:"carModelId"`
	IsMainImage bool            `json:"isMainImage"`
	Image       GetFileResponse `json:"image"`
}

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required,max=100"`
}
type UpdateCarModelPropertyRequest struct {
	Value string `json:"value" binding:"max=100"`
}

type CarModelPropertyResponse struct {
	Id         int                 `json:"id"`
	CarModelId int                 `json:"carModelId"`
	PropertyId int                 `json:"propertyId"`
	Property   GetPropertyResponse `json:"property"`
	Value      string              `json:"value"`
}

type CreateCarModelCommentRequest struct {
	UserId     int    `json:"userId"`
	CarModelId int    `json:"carModelId" binding:"required"`
	Message    string `json:"message" binding:"required,max=1000"`
}
type UpdateCarModelCommentRequest struct {
	Message string `json:"message" binding:"required,max=1000"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	UserId     int          `json:"userId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}
