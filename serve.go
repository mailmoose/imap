package imapbackend

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/emersion/go-imap/server"
)

func Serve(config *Config) {

	log.SetLevel(log.DebugLevel)

	err := InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v")
	}

	if config.SeedDB {
		seedDB()
	}

	// Create the backend
	backend, err := New("")
	if err != nil {
		log.Fatalf("Couldn't create backend: %v")
	}
	_ = backend

	// Create a IMAP new server
	s := server.New(backend)
	s.Addr = config.IMAPAddress
	// Since we will use this server for testing only, we can allow plain text
	// authentication over unencrypted connections
	s.AllowInsecureAuth = true

	s.Debug = os.Stderr

	log.Printf("Starting IMAP server at %s", config.IMAPAddress)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
