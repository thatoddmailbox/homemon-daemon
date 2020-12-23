package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	// Interval sets the amount of time, in seconds, to wait between reports.
	Interval int

	// Host sets the host to submit reports to.
	Host string

	// Token sets the token used to authenticate reports.
	Token string

	// Transport sets the transport to use to send reports.
	Transport string
}

var currentConfig config

func loadConfig() error {
	_, err := toml.DecodeFile("config.toml", &currentConfig)
	if os.IsNotExist(err) {
		log.Fatalln("Could not find config file.")
	}
	return err
}
