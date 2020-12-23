package main

import (
	"io/ioutil"
	"log"
	"strings"
)

const baseURL = "https://homemon-rpt.studer.dev/"
const reportURL = baseURL + "report"

func report(t transport) error {
	tokenBytes, err := ioutil.ReadFile("token.txt")
	if err != nil {
		return err
	}
	tokenString := strings.TrimSpace(string(tokenBytes))

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

	return t.Transport([]byte(tokenString), powered, batteryCapacity, batteryVoltage)
}

func main() {
	log.Println("homemon-daemon")

	err := report(&transportHTTP{})
	if err != nil {
		panic(err)
	}
}
