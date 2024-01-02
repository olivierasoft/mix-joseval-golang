package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/domain/entity"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/domain/exception"
	"github.com/olivierasoft/mix-joseval-golang.git/internal/domain/res"
)

func getEncodedParamsBuffer(code string) *bytes.Buffer {
	formData := url.Values{}
	formData.Set("client_id", os.Getenv("DISCORD_CLIENT_ID"))
	formData.Set("client_secret", os.Getenv("DISCORD_SECRET"))
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", os.Getenv("DISCORD_REDIRECT_URL"))
	formData.Set("scope", "identity")

	return bytes.NewBufferString(formData.Encode())
}

func getUserAuthorization(code string) (*http.Response, error) {

	body := getEncodedParamsBuffer(code)

	tokenRequest, err := http.NewRequest(http.MethodPost,
		"https://discord.com/api/oauth2/token",
		body,
	)

	if err != nil {
		return nil, err
	}

	tokenRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return HttpClient.Do(tokenRequest)

}

func retrieveUserInformation(token string) (*res.RetrieveUserInformationResponse, error) {
	tokenRequest, err := http.NewRequest(http.MethodGet, "https://discord.com/api/oauth2/@me", nil)

	if err != nil {
		return nil, err
	}

	tokenRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	response, requestErr := HttpClient.Do(tokenRequest)

	if requestErr != nil {
		return nil, err
	}

	bytes, readErr := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		var discordError exception.DiscordErrorRequest

		fmt.Println(string(bytes))

		json.Unmarshal(bytes, &discordError)

		return nil, fmt.Errorf(discordError.Error)
	}

	if readErr != nil {
		return nil, err
	}

	var retrieveUserInformationResponse res.RetrieveUserInformationResponse

	marshalErr := json.Unmarshal(bytes, &retrieveUserInformationResponse)

	if marshalErr != nil {
		return nil, err
	}

	return &retrieveUserInformationResponse, nil
}

func Login(code string) (token *string, err error) {

	response, err := getUserAuthorization(code)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		var discordError exception.DiscordErrorRequest

		json.Unmarshal(bodyBytes, &discordError)

		return nil, fmt.Errorf(discordError.ErrorDescription)
	}

	var authRequest res.DiscordAuthRequest

	json.Unmarshal(bodyBytes, &authRequest)

	userInformation, err := retrieveUserInformation(authRequest.AccessToken)

	if err != nil {
		return nil, err
	}

	user := entity.User{
		Username:      userInformation.User.Username,
		Avatar:        userInformation.User.Avatar,
		GlobalName:    userInformation.User.GlobalName,
		Discriminator: userInformation.User.Discriminator,
	}

	tx := DatabaseClient.Create(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Name":      user.GlobalName,
		"Username":  user.Username,
		"ExpiresAt": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		"IssuedAt":  jwt.NewNumericDate(time.Now()),
	})

	signedString, err := jwt.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}

	return &signedString, nil
}
