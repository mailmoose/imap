package imap

import (
	"fmt"
)

var (
	defaultIMAPPort    = 143
	defaultIMAPAddress = fmt.Sprintf(":%d", defaultIMAPPort)
)

// Config contains all the config for serving the IMAP backend
type Config struct {

	// IMAPAddress is the address on which the IMAP server will be exposed.
	IMAPAddress string

	// TlsCert is the public certificate.
	TlsCert string

	// TlsKey is the certificates private key.
	TlsKey string

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
