package schemas

type Permission struct {
	Id   int64 `gorm:"primaryKey;autoIncrement:true;"`
	Name string
}
