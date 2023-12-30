package business

import (
	"net/http"
)

type RequestCredentials struct {
	Username string
	Password string
}

func RetrieveDiscordUserBearerToken(code string) (*http.Response, error) {

	tokenUrl := "https://discord.com/oauth2/token"

	http.NewRequest(http.MethodPost)

	if err != nil {
		return nil, err
	}

	return response, nil
}
