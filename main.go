package main

import (
	"log"
	"os/exec"
	"runtime"
	"time"
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

	var t transport
	if currentConfig.Transport == "HTTP" {
		t = &transportHTTP{}
	} else if currentConfig.Transport == "UDP" {
		t = &transportUDP{}
	} else {
		log.Fatalf("Unknown transport '%s'.", currentConfig.Transport)
	}

	if currentConfig.InitialDelay.Duration > 0 {
		time.Sleep(currentConfig.InitialDelay.Duration)
	}

	count := 0
	for {
		err = report([]byte(currentConfig.Token), t)
		if err != nil {
			log.Println("Error while sending report!")
			log.Println(err)
		}

		lastTime := time.Now()
		runtime.GC()
		sleepTime := currentConfig.Interval.Duration - time.Since(lastTime)
		time.Sleep(sleepTime)

		if currentConfig.RestartCount != 0 {
			count++
			if count == currentConfig.RestartCount {
				err := exec.Command("reboot").Run()
				if err != nil {
					log.Println("Error while rebooting!")
					log.Println(err)
				}
			}
		}
	}
}
