package entity

import (
	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	Id       int32
	Platform string `gorm:"not null;size:140"`
}
