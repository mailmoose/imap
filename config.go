package imap

import (
	"crypto/tls"
	"fmt"
)

var (
	DefaultIMAPPort    = 143
	DefaultIMAPAddress = fmt.Sprintf(":%d", DefaultIMAPPort)
)

// Config contains all the config for serving the IMAP backend
type Config struct {

	// IMAPAddress is the address on which the IMAP server will be exposed.
	IMAPAddress string

	// TLSConfig for the IMAP server
	TLSConfig *tls.Config

	// Debug flag.
	Debug bool
}

// Validate validates whether all config is set and valid
func (config *Config) Validate() error {

	if config.IMAPAddress == "" {
		return fmt.Errorf("IMAPAddress cannot be empty")
	}

	return nil
}
