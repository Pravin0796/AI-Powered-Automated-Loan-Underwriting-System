package experian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// OAuthTokenResponse represents the response from Experian OAuth
type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// GetAccessToken retrieves an OAuth token from Experian
func GetAccessToken() (string, error) {
	LoadEnv()
	oauthURL := "https://sandbox-us-api.experian.com/oauth2/v1/token"
	
	reqBody, _ := json.Marshal(map[string]string{
		"client_id":     os.Getenv("EXPERIAN_CLIENT_ID"),
		"client_secret": os.Getenv("EXPERIAN_SECRET"),
		"grant_type":    "client_credentials",
	})

	req, err := http.NewRequest("POST", oauthURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var tokenResp OAuthTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}
