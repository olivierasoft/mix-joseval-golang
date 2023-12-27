package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        int64  `gorm:"primaryKey;autoIncrement:true"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsActive  bool   `json:"isActive"`
	Roles     []Role `gorm:"many2many:user_role"`
}
