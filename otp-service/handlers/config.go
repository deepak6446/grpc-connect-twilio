package handlers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// TwilioAPIConfig holds Twilio API credentials
type TwilioAPIConfig struct {
	AccountSID string
	AuthToken  string
	FromNumber string
}

// GetTwilioAPIConfig retrieves Twilio API credentials from environment variables
func GetTwilioAPIConfig() *TwilioAPIConfig {
	return &TwilioAPIConfig{
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromNumber: os.Getenv("TWILIO_FROM_NUMBER"),
	}
}
