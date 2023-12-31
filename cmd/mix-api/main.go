package main

import (
	"github.com/olivierasoft/mix-joseval-golang.git/config"
)

func init() {
	config.ReadEnvFile()
}

func main() {
	db, err := config.GenDatabaseInstance()

	if err != nil {
		panic(err)
	}

	config.AddMigrations(db)

	config.BootstrapGinFramework()
}
