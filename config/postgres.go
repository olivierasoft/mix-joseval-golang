package config

import (
	"github.com/olivierasoft/mix-joseval-golang.git/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgresql://postgres:@@mix@@@localhost:5432/mix"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schemas.User{}, &schemas.Permission{}, &schemas.Role{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
