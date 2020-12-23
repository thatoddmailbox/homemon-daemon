package main

type transport interface {
	Transport(token []byte, powered usbStatus, batteryCapacity uint8, batteryVoltage uint16) error
}
