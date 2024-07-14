package models

import "time"

type City struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}

type Country struct {
	BaseModel
	Name   string `gorm:"type:string;size:15;not null"`
	Cities []City
}

type PersianYear struct {
	Title   string    `gorm:"type:string;size:10;not null;unique"`
	Year    int       `gorm:"type:int;not null;uniqueIndex"`
	StartAt time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt   time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
}
