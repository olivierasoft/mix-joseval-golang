package entity

import "gorm.io/gorm"

type DiscordConnection struct {
	gorm.Model
	Id         int32
	Name       string `gorm:"size:256"`
	App        string `gorm:"size:120"`
	FriendSync bool
	Verified   bool
	UserID     int32
}
