package config

import (
	"os"
	"strings"

	"github.com/olivierasoft/mix-joseval-golang.git/schemas"
)

func ReadEnvironmentFile() (err error) {
	bytes, err := os.ReadFile(".env")

	if err != nil {
		panic(err)
	}

	environmentSplit := strings.Split(string(bytes), "=")

	if len(environmentSplit)%2 != 0 {
		return &schemas.Error{
			Message: "Environment file in wrong format.",
		}
	}

	for i := 0; i < len(environmentSplit); i += 2 {
		os.Setenv(environmentSplit[i], environmentSplit[i+1])
	}

	return nil
}
