package imapbackend

import (
	"fmt"
	"os"
)

var (
	defaultIMAPPort    = 143
	defaultIMAPAddress = fmt.Sprintf(":%d", defaultIMAPPort)
	defaultDatabaseURL = "sqlite:test.db"
	defaultSeedDB      = false
)

// BuildConfigFromEnv populates a IMAP backend config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.IMAPAddress = getEnv("IMAP_ADDRESS", defaultIMAPAddress)
	config.DatabaseURL = getEnv("DATABASE_URL", defaultDatabaseURL)

	seedDB := getEnv("SEED_DB", "0")
	if seedDB == "1" {
		config.SeedDB = true
	} else {
		config.SeedDB = false
	}

	return config
}

// Config contains all the config for serving the IMAP backend
type Config struct {
	IMAPAddress string
	DatabaseURL string
	SeedDB      bool
}

// Validate validates whether all config is set and valid
func (config *Config) Validate() error {

	if config.IMAPAddress == "" {
		return fmt.Errorf("IMAPAddress cannot be empty")
	}

	if config.DatabaseURL == "" {
		return fmt.Errorf("DatabaseURL cannot be empty")
	}

	return nil
}

// getEnv gets the env variable with the given key if the key exists
// else it falls back to the fallback value
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
