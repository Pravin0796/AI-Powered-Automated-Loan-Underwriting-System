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
	fmt.Println("Fetching Credit Report...")
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

	fmt.Println(apiURL)
	// Example request payload (modify as needed)
	data := map[string]interface{}{
		"consumerPii": map[string]interface{}{
			"primaryApplicant": map[string]interface{}{
				"name": map[string]string{
					"firstName": "John",
					"lastName":  "Doe",
				},
				"dob": map[string]string{
					"dob": "1990-01-01",
				},
				"ssn": map[string]string{
					"ssn": "666-00-0001",
				},
				"currentAddress": map[string]string{
					"line1":   "123 Main Street",
					"city":    "Costa Mesa",
					"state":   "CA",
					"zipCode": "92626",
					"country": "US",
				},
			},
		},
		"requestor": map[string]string{
			"subscriberCode": "1234567",
		},
		"permissiblePurpose": map[string]string{
			"type": "01",
		},
	}

	reqBody, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	//req.Header.Set("companyId", os.Getenv("EXPERIAN_COMPANY_ID"))
	req.Header.Set("clientReferenceId", "SBMYSQL")

	//fmt.Printf("Request:\nURL: %s\nHeaders: %+v\nBody: %s\n", req.URL.String(), req.Header, string(reqBody))

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
