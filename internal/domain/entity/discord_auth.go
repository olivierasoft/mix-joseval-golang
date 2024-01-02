package entity

import (
	"time"

	"gorm.io/gorm"
)

type DiscordAuth struct {
	gorm.Model
	Id             int32
	Token          string    `gorm:"not null;size:256"`
	RefreshToken   string    `gorm:"not null;size:256"`
	ExpirationDate time.Time `gorm:"not null"`
	Active         bool      `gorm:"default:true"`
	UserID         int32
}
