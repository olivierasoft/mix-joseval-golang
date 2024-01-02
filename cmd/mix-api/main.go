package main

import (
	"github.com/olivierasoft/mix-joseval-golang.git/config"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/app/service"
)

func init() {
	config.ReadEnvFile()

}

func main() {
	db, err := config.GenDatabaseInstance()

	if err != nil {
		panic(err)
	}

	service.SetHttpGlobalClient(config.GetHttpClient())
	service.SetGlobalDatabaseClient(db)

	config.AddMigrations(db)
	config.BootstrapGinFramework()
}
