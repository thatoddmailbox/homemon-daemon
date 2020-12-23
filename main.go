package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var host string

func report(token []byte, t transport) error {
	usbPresent, err := getUSBPresent()
	powered := usbStatusNotPresent
	if err != nil {
		log.Println("Error getting USB status!")
		log.Println(err)
		powered = usbStatusError
	} else if usbPresent {
		powered = usbStatusPresent
	}

	batteryCapacity, err := getBatteryCapacity()
	if err != nil {
		log.Println("Error getting battery capacity!")
		log.Println(err)
	}

	batteryVoltage, err := getBatteryVoltage()
	if err != nil {
		log.Println("Error getting battery voltage!")
		log.Println(err)
	}

	return t.Transport(token, powered, batteryCapacity, batteryVoltage)
}

func main() {
	log.Println("homemon-daemon")

	tokenBytes, err := ioutil.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	tokenBytes = []byte(strings.TrimSpace(string(tokenBytes)))

	hostBytes, err := ioutil.ReadFile("host.txt")
	if err != nil {
		panic(err)
	}
	host = strings.TrimSpace(string(hostBytes))

	err = report(tokenBytes, &transportHTTP{})
	if err != nil {
		panic(err)
	}
}
