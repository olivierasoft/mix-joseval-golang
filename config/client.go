package config

import (
	"net/http"
	"time"
)

func GetHttpClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
