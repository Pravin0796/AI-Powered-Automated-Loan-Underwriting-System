package experian

import (
	"encoding/json"
	"fmt"
	"log"

	"AI-Powered-Automated-Loan-Underwriting-System/config"
	"github.com/go-resty/resty/v2"
)

// CreditProfileRequest defines the request body structure
type CreditProfileRequest struct {
	SSN       string `json:"ssn"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Address   string `json:"address"`
	City      string `json:"city"`
	State     string `json:"state"`
	ZipCode   string `json:"zipCode"`
}

// CreditProfileResponse defines the API response structure
type CreditProfileResponse struct {
	CreditScore  int    `json:"creditScore"`
	CreditStatus string `json:"creditStatus"`
}

// FetchCreditProfile calls Experian's API
func FetchCreditProfile(ssn, lastName, firstName, address, city, state, zip string) (*CreditProfileResponse, error) {
	client := resty.New()

	// Create the request body
	requestBody := CreditProfileRequest{
		SSN:       ssn,
		LastName:  lastName,
		FirstName: firstName,
		Address:   address,
		City:      city,
		State:     state,
		ZipCode:   zip,
	}

	// Send API request
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Basic "+config.ExperianSecret).
		//SetHeader("x-api-key", config.ExperianAPIKey).
		SetBody(requestBody).
		Post(config.ExperianAPIBaseURL + "/v2/credit-report") // Use full API URL

	if err != nil {
		log.Println("Error calling Experian API:", err)
		return nil, err
	}

	// Handle API response
	if resp.IsError() {
		return nil, fmt.Errorf("Experian API error: %s", resp.String())
	}

	// Parse JSON response
	var creditProfile CreditProfileResponse
	err = json.Unmarshal(resp.Body(), &creditProfile)
	if err != nil {
		return nil, err
	}

	return &creditProfile, nil
}
