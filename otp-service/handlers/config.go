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
		log.Printf("Error loading .env file: ", err)
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
		AccountSID: getEnv("TWILIO_ACCOUNT_SID"),
		AuthToken:  getEnv("TWILIO_AUTH_TOKEN"),
		FromNumber: getEnv("TWILIO_FROM_NUMBER"),
	}
}

// getEnv retrieves an environment variable or returns an empty string
func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Environment variable %s not found", key)
	}
	return val
}
