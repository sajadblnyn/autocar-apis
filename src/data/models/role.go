package models

type Role struct {
	BaseModel
	Name      string `gorm:"type:string;not null;unique;size:10"`
	UserRoles *[]UserRole
}
