package main

import (
	"github.com/olivierasoft/mix-joseval-golang.git/config"
	"github.com/olivierasoft/mix-joseval-golang.git/router"
)

func init() {
	error := config.Init()

	if error != nil {
		panic(error)
	}
}

func main() {
	router.Initialize()
}
