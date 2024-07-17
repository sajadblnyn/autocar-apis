package models

import "time"

type City struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
}

type Country struct {
	BaseModel
	Name      string `gorm:"type:string;size:15;not null"`
	Cities    []City
	Companies []Company
}

type PersianYear struct {
	BaseModel
	PersianTitle  string    `gorm:"type:string;size:10;not null;unique"`
	Year          int       `gorm:"type:int;not null;uniqueIndex"`
	StartAt       time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt         time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	CarModelYears []CarModelYear
}

type Color struct {
	BaseModel
	Name           string `gorm:"type:string;not null;unique;size:15;"`
	HexCode        string `gorm:"type:string;not null;unique;size:7;"`
	CarModelColors []CarModelColor
}

type File struct {
	BaseModel
	Name        string `gorm:"type:string;not null;size:100;"`
	Directory   string `gorm:"type:string;not null;size:100;"`
	Description string `gorm:"type:string;not null;size:500;"`
	MediaType   string `gorm:"type:string;not null;size:20;"`
}
