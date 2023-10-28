package imap

import (
	log "github.com/sirupsen/logrus"

	"github.com/emersion/go-imap/backend"
	"github.com/emersion/go-imap/server"
)

func Serve(config *Config, backend backend.Backend) {

	log.SetLevel(log.DebugLevel)

	// Create a IMAP new server
	s := server.New(backend)
	s.Addr = config.IMAPAddress
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = true

	// Log with logrus
	logrusLogger := log.New()
	logWriter := logrusLogger.Writer()
	defer logWriter.Close()

	if config.Debug {
		s.Debug = logWriter
	}
	s.ErrorLog = logrusLogger

	// TLS config
	s.TLSConfig = config.TLSConfig

	log.Printf("Starting IMAP server at %s", config.IMAPAddress)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
