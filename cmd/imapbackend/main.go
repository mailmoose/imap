package main

import (
	log "github.com/sirupsen/logrus"

	imapbackend "github.com/gopistolet/gopistolet/backend/imap"
	imap "github.com/gopistolet/imap"
)

func main() {

	config := imap.BuildConfigFromEnv()

	err := config.Validate()
	if err != nil {
		log.Fatalf("config file invalid: %v", err)
	}

	// Create backends
	db, err := imap.InitDB(config.DatabaseURL)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}

	backendIMAP, err := imapbackend.NewIMAPBackend(db)
	if err != nil {
		log.Fatalf("Couldn't create IMAP backend: %v", err)
	}

	log.Printf("starting IMAP backend with config: %+v", config)

	imap.Serve(config, backendIMAP)

}
