package experian

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FetchCreditReport requests a credit report from Experian
func FetchCreditReport() {
	err := godotenv.Load("../.env") // Adjust for cmd directory
	if err != nil {
		log.Panic("Error loading .env file:", err)
	}
	baseURL := os.Getenv("EXPERIAN_API_URL")
	accessToken, err := GetAccessToken()
	if err != nil {
		fmt.Println("Error fetching access token:", err)
		return
	}

	apiURL := baseURL + "/v2/credit-report"

	// Example request payload (modify as needed)
	reqBody, _ := json.Marshal(map[string]interface{}{
		"firstName": "John",
		"lastName":  "Doe",
		"ssn":       "123-45-6789", // Replace with test data
		"dob":       "1990-01-01",
	})

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Credit Report Response:", string(body))
}
