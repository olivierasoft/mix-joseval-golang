package schemas

type Role struct {
	Id          int64 `gorm:"primaryKey;autoIncrement:true"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permission"`
}
