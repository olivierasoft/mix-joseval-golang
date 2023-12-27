package config

import "gorm.io/gorm"

var db *gorm.DB

func Init() error {

	postgres, err := InitializePostgres()

	if err != nil {
		return err
	}

	db = postgres

	return nil
}
