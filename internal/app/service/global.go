package service

import (
	"net/http"

	"gorm.io/gorm"
)

var HttpClient *http.Client
var DatabaseClient *gorm.DB

func SetGlobalDatabaseClient(gorm *gorm.DB) {
	DatabaseClient = gorm
}

func SetHttpGlobalClient(client *http.Client) {
	HttpClient = client
}
