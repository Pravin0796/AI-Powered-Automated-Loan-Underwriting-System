package config

import (
	"os"
)

var (
	ExperianAPIBaseURL = os.Getenv("EXPERIAN_API_URL")
	ExperianClientID   = os.Getenv("EXPERIAN_CLIENT_ID")
	ExperianSecret     = os.Getenv("EXPERIAN_SECRET")
	ExperianAPIKey     = os.Getenv("EXPERIAN_API_KEY")
)
