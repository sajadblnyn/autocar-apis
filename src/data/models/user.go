package models

type User struct {
	BaseModel
	Username     string `gorm:"type:string;not null;unique;size:20"`
	Password     string `gorm:"type:string;not null;size:64;"`
	FirstName    string `gorm:"type:string;null;size:15"`
	LastName     string `gorm:"type:string;null;size:25"`
	MobileNumber string `gorm:"type:string;null;size:11;default:null;unique"`
	Email        string `gorm:"type:string;null;size:64;default:null;unique"`
	Enabled      bool   `gorm:"default:true"`
	UserRoles    *[]UserRole
}
