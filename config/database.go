package config

import (
	"os"

	"github.com/olivierasoft/mix-joseval-golang.git/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GenDatabaseInstance() (instance *gorm.DB, err error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddMigrations(db *gorm.DB) {
	db.AutoMigrate(&entity.User{}, &entity.DiscordConnection{}, &entity.DiscordGuild{})
}
