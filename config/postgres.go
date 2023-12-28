package config

import (
	"os"

	"github.com/olivierasoft/mix-joseval-golang.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {

	DATABASE_URL := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(DATABASE_URL), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Permission{}, &schemas.Role{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
