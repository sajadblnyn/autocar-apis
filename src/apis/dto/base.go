package dto

import (
	"mime/multipart"
	"time"
)

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,max=20,min=3"`
}

type GetCountryResponse struct {
	Id        int               `json:"id"`
	Name      string            `json:"name"`
	Cities    []GetCityResponse `json:"cities,omitempty"`
	Companies []CompanyResponse `json:"companies,omitempty"`
}

type CreateUpdateCityRequest struct {
	Name      string `json:"name" binding:"required,alpha,max=20,min=3"`
	CountryId int    `json:"countryId,omitempty"`
}

type GetCityResponse struct {
	Id      int                `json:"id"`
	Name    string             `json:"name"`
	Country GetCountryResponse `json:"country"`
}

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" binding:"required"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Directory   string `json:"directory" binding:"required"`
	MediaType   string `json:"mediaType" binding:"required"`
}

type UpdateFileRequest struct {
	Description string `json:"description" binding:"required"`
}

type GetFileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Directory   string `json:"directory"`
	MediaType   string `json:"mediaType"`
}

type CreateColorRequest struct {
	Name    string `json:"name" binding:"required,alpha,max=15,min=3"`
	HexCode string `json:"hexCode" binding:"required,max=7,min=7"`
}
type UpdateColorRequest struct {
	Name    string `json:"name,omitempty" binding:"alpha,max=15,min=3"`
	HexCode string `json:"hexCode,omitempty" binding:"max=7,min=7"`
}

type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

type CreatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle" binding:"required,max=4,min=4"`
	Year         int       `json:"year" binding:"required"`
	StartAt      time.Time `json:"startAt" binding:"required"`
	EndAt        time.Time `json:"endAt" binding:"required"`
}

type UpdatePersianYearRequest struct {
	PersianTitle string    `json:"persianTitle,omitempty" binding:"max=4,min=4"`
	Year         int       `json:"year,omitempty"`
	StartAt      time.Time `json:"startAt,omitempty"`
	EndAt        time.Time `json:"endAt,omitempty"`
}

type PersianYearResponse struct {
	Id           int       `json:"id"`
	PersianTitle string    `json:"persianTitle"`
	Year         int       `json:"year"`
	StartAt      time.Time `json:"startAt"`
	EndAt        time.Time `json:"endAt"`
}

type PersianYearWithoutDateResponse struct {
	Id           int    `json:"id"`
	PersianTitle string `json:"persianTitle"`
}
