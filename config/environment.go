package config

import (
	"bufio"
	"os"
	"strings"

	"github.com/olivierasoft/mix-joseval-golang.git/schemas"
)

func ReadEnvironmentFile() (err error) {
	file, err := os.Open(".env")

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		environmentSplit := strings.Split(fileScanner.Text(), "=")

		if len(environmentSplit)%2 != 0 {
			return &schemas.Error{
				Message: "Environment file in wrong format.",
			}
		}

		for i := 0; i < len(environmentSplit); i += 2 {
			os.Setenv(environmentSplit[i], environmentSplit[i+1])
		}

	}

	return nil
}
