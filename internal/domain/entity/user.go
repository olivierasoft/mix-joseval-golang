package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id                 int32
	Username           string    `gorm:"not null;size:256"`
	Avatar             string    `gorm:"size:256"`
	Discriminator      string    `gorm:"size:40"`
	GlobalName         string    `gorm:"size:150"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	DiscordGuilds      []DiscordGuild
	DiscordConnections []DiscordConnection
}
