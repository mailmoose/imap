package imap

import (
	"crypto/tls"

	imapbackend "github.com/mistralmail/mistralmail/backend/imap"
	log "github.com/sirupsen/logrus"

	"github.com/emersion/go-imap/server"
)

func Serve(config *Config, backend *imapbackend.IMAPBackend) {

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
	if config.TlsCert != "" && config.TlsKey != "" {
		cert, err := tls.LoadX509KeyPair(config.TlsCert, config.TlsKey)
		if err != nil {
			log.Fatalf("Could not load keypair: %v", err)
		} else {
			s.TLSConfig = &tls.Config{
				Certificates: []tls.Certificate{cert},
			}
		}
	}

	log.Printf("Starting IMAP server at %s", config.IMAPAddress)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
