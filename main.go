package main

import (
	"log"
)

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

	err := loadConfig()
	if err != nil {
		panic(err)
	}

	err = report([]byte(currentConfig.Token), &transportHTTP{})
	if err != nil {
		panic(err)
	}
}
