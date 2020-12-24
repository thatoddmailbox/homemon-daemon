package main

import (
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

type config struct {
	// Interval sets the amount of time to wait between reports.
	Interval duration

	// Host sets the host to submit reports to.
	Host string

	// Port sets the port on the Host to submit reports to. (for UDP only!)
	Port int

	// Token sets the token used to authenticate reports.
	Token      string
	tokenBytes []byte

	// Transport sets the transport to use to send reports.
	Transport string
}

var currentConfig config

func loadConfig() error {
	_, err := toml.DecodeFile("config.toml", &currentConfig)
	if os.IsNotExist(err) {
		log.Fatalln("Could not find config file.")
	} else if err != nil {
		return err
	}

	currentConfig.tokenBytes = []byte(currentConfig.Token)

	if currentConfig.Transport == "UDP" {
		currentConfig.tokenBytes, err = base64.URLEncoding.DecodeString(currentConfig.Token)
		if err != nil {
			return err
		}

		if len(currentConfig.tokenBytes) != 64 {
			log.Fatalf("For UDP transport, token must be exactly 64 bytes long.")
		}
	}

	return nil
}
