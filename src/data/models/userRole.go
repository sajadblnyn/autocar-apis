package models

type UserRole struct {
	BaseModel
	UserId int
	RoleId int
	User   User `gorm:"foreignKey:UserId"`
	Role   Role `gorm:"foreignKey:RoleId"`
}
