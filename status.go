package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const powerSupplyFolder = "/sys/class/power_supply/"

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

func getBatteryCapacity() (int, error) {
	return readPowerSupplyFile("battery/capacity")
}

func getBatteryVoltage() (int, error) {
	v, err := readPowerSupplyFile("battery/voltage_now")
	if v != -1 {
		v = v / 1000
	}
	return v, err
}
