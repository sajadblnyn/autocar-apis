package models

type City struct {
	BaseModel
	Name      string `gorm:"type:string;size:10;not null"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId"`
}
