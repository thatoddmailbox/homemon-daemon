package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const powerSupplyFolder = "/sys/class/power_supply/"

type usbStatus int8

const (
	usbStatusNotPresent usbStatus = iota
	usbStatusPresent
	usbStatusError
)

func readPowerSupplyFile(filename string) (int, error) {
	data, err := ioutil.ReadFile(powerSupplyFolder + filename)
	if err != nil {
		return -1, err
	}

	dataString := strings.TrimSpace(string(data))

	dataInt, err := strconv.Atoi(dataString)
	if err != nil {
		return -1, err
	}

	return dataInt, nil
}

func getBatteryCapacity() (uint8, error) {
	c, err := readPowerSupplyFile("battery/capacity")
	return uint8(c), err
}

func getBatteryVoltage() (uint16, error) {
	v, err := readPowerSupplyFile("battery/voltage_now")
	if v != -1 {
		v = v / 1000
	}
	return uint16(v), err
}

func getUSBPresent() (bool, error) {
	p, err := readPowerSupplyFile("usb/present")
	if err != nil {
		if os.IsNotExist(err) {
			// that's ok, that probably just means there's no usb host
			// (aka, we're not connected to power)
			return false, nil
		}

		return false, err
	}

	return p == 1, nil
}
