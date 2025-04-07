package experian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

// OAuthTokenResponse represents the response from Experian OAuth
type OAuthTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// GetAccessToken retrieves an OAuth token from Experian
//func GetAccessToken() (string, error) {
//	LoadEnv()
//	oauthURL := "https://sandbox-us-api.experian.com/oauth2/v1/token"
//
//	//reqBody, _ := json.Marshal(map[string]string{
//	//	"client_id":     os.Getenv("EXPERIAN_CLIENT_ID"),
//	//	"client_secret": os.Getenv("EXPERIAN_SECRET"),
//	//	"grant_type":    "client_credentials",
//	//})
//	//
//	//req, err := http.NewRequest("POST", oauthURL, bytes.NewBuffer(reqBody))
//	//if err != nil {
//	//	return "", err
//	//}
//	//
//	//req.Header.Set("Content-Type", "application/json")
//
//	data := "client_id=" + os.Getenv("EXPERIAN_CLIENT_ID") +
//		"&client_secret=" + os.Getenv("EXPERIAN_SECRET") +
//		"&grant_type=client_credentials"
//
//	req, err := http.NewRequest("POST", oauthURL, bytes.NewBufferString(data))
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//	if resp.StatusCode != http.StatusOK {
//		fmt.Printf("Error response from Experian: %d\nBody: %s\n", resp.StatusCode, body)
//		return "", fmt.Errorf("failed to get access token")
//	}
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	body, _ := ioutil.ReadAll(resp.Body)
//	var tokenResp OAuthTokenResponse
//	if err := json.Unmarshal(body, &tokenResp); err != nil {
//		return "", err
//	}
//
//	fmt.Println(tokenResp.AccessToken)
//	return tokenResp.AccessToken, nil
//}

func GetAccessToken() (string, error) {
	fmt.Println("Calling GetAccessToken...")

	LoadEnv()
	clientID := os.Getenv("EXPERIAN_CLIENT_ID")
	clientSecret := os.Getenv("EXPERIAN_SECRET")
	clientUsername := os.Getenv("EXPERIAN_CLIENT_USERNAME")
	clientPassword := os.Getenv("EXPERIAN_CLIENT_PASSWORD")

	if clientID == "" || clientSecret == "" {
		return "", fmt.Errorf("Missing client credentials")
	}

	oauthURL := "https://sandbox-us-api.experian.com/oauth2/v1/token"

	payload := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"username":      clientUsername,
		"password":      clientPassword,
		//"grant_type":    "client_credentials",
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

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
	//fmt.Println("Response status:", resp.StatusCode)
	//fmt.Println("Response body:", string(body))

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get access token: %s", string(body))
	}

	// Parse JSON manually
	var tokenMap map[string]interface{}
	if err := json.Unmarshal(body, &tokenMap); err != nil {
		return "", err
	}

	accessToken, ok := tokenMap["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("access_token not found in response")
	}

	fmt.Println("Access Token:", accessToken)
	return accessToken, nil
}
