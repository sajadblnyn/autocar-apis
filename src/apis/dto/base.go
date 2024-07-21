package dto

import "mime/multipart"

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
