package imap

import (
	"fmt"
	"os"
)

var (
	defaultIMAPPort    = 143
	defaultIMAPAddress = fmt.Sprintf(":%d", defaultIMAPPort)
	defaultDatabaseURL = "sqlite:test.db"
)

// BuildConfigFromEnv populates a IMAP backend config from env variables
func BuildConfigFromEnv() *Config {
	config := &Config{}

	config.IMAPAddress = getEnv("IMAP_ADDRESS", defaultIMAPAddress)
	config.DatabaseURL = getEnv("DATABASE_URL", defaultDatabaseURL)

	config.TlsCert = getEnv("TLS_CERT", "")
	config.TlsKey = getEnv("TLS_KEY", "")

	return config
}

// Config contains all the config for serving the IMAP backend
type Config struct {
	IMAPAddress string
	DatabaseURL string
	SeedDB      bool
	TlsCert     string
	TlsKey      string
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
