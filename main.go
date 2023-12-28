package main

import (
	"fmt"
	"os"

	"github.com/olivierasoft/mix-joseval-golang.git/config"
	"github.com/olivierasoft/mix-joseval-golang.git/router"
)

func init() {

	envErr := config.ReadEnvironmentFile()

	if envErr != nil {
		panic(envErr)
	}

	configErr := config.Init()

	if configErr != nil {
		panic(configErr)
	}
}

func main() {

	fmt.Println(os.Getenv("DISCORD_SECRET"))

	router.Initialize()
}
