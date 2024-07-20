package models

type PropertyCategory struct {
	BaseModel
	Name       string     `gorm:"size:50;type:string;not null;unique"`
	Icon       string     `gorm:"size:1000;type:string;not null;"`
	Properties []Property `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
}

type Property struct {
	BaseModel
	Name        string `gorm:"size:50;type:string;not null;unique"`
	Icon        string `gorm:"size:1000;type:string;not null;"`
	CategoryId  int
	Category    PropertyCategory `gorm:"foreignKey:CategoryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION;"`
	Description string           `gorm:"size:1000;type:string;not null;"`
	DataType    string           `gorm:"size:15;type:string;not null;"`
	Unit        string           `gorm:"size:15;type:string;not null;"`
}
