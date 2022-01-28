package main

import (
	log "github.com/sirupsen/logrus"

	imapbackend "github.com/gopistolet/imap-backend"
)

func main() {

	config := imapbackend.BuildConfigFromEnv()

	err := config.Validate()
	if err != nil {
		log.Fatalf("config file invalid: %v", err)
	}

	log.Printf("starting IMAP backend with config: %+v", config)

	imapbackend.Serve(config)

}
