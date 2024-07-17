package models

import "time"

type Gearbox struct {
	BaseModel
	Name     string `gorm:"type:string;size:15;not null;unique;"`
	CarModel []CarModel
}

type CarType struct {
	BaseModel
	Name     string `gorm:"type:string;size:15;not null;unique;"`
	CarModel []CarModel
}

type Company struct {
	BaseModel
	Name      string `gorm:"type:string;size:15;not null;unique;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModel  []CarModel
}

type CarModel struct {
	BaseModel
	Name               string `gorm:"type:string;size:15;not null;unique;"`
	CompanyId          int
	Company            Company `gorm:"foreignKey:CompanyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarTypeId          int
	CarType            CarType `gorm:"foreignKey:CarTypeId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	GearboxId          int
	Gearbox            Gearbox `gorm:"foreignKey:GearboxId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelColors     []CarModelColor
	CarModelYears      []CarModelYear
	CarModelProperties []CarModelProperty
	CarModelImages     []CarModelImage
	CarModelComments   []CarModelComment
}

type CarModelColor struct {
	BaseModel
	Color      Color `gorm:"foreignKey:ColorId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	ColorId    int
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelId int
}

type CarModelYear struct {
	BaseModel
	CarModel               CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelId             int
	PersianYear            PersianYear `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	PersianYearId          int
	CarModelPriceHistories []CarModelPriceHistory
}

type CarModelImage struct {
	BaseModel
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelId  int
	Image       File `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	ImageId     int
	IsMainImage bool
}

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelId int
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	PropertyId int
	Value      string `gorm:"size:100;type:string;not null"`
}

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	CarModelId int
	User       User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	UserId     int
	Message    string `gorm:"size:1000;type:string;not null"`
}
