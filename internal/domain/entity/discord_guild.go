package entity

import "gorm.io/gorm"

type DiscordGuild struct {
	gorm.Model
	Id     int32
	Name   string `gorm:"size:256"`
	Icon   string `gorm:"size:256"`
	Owner  bool
	UserID int32
}
