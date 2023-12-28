package main

import (
	"os"

	"github.com/olivierasoft/mix-joseval-golang.git/config"
	"github.com/olivierasoft/mix-joseval-golang.git/router"
)

func init() {

	isProduction := os.Getenv("production")

	if isProduction == "" {
		envErr := config.ReadEnvironmentFile()

		if envErr != nil {
			panic(envErr)
		}
	}

	configErr := config.Init()

	if configErr != nil {
		panic(configErr)
	}
}

func main() {
	router.Initialize()
}
