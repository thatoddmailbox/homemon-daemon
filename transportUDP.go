package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"net"
	"strconv"
	"time"
)

const batteryCapacityError = 0x7F
const batteryVoltageError = 0x1FFF

type transportUDP struct{}

func (t *transportUDP) calculateHMAC(message []byte) ([]byte, error) {
	mac := hmac.New(sha256.New, currentConfig.tokenBytes)
	_, err := mac.Write(message)
	if err != nil {
		return nil, err
	}
	return mac.Sum(nil), nil
}

func (t *transportUDP) Transport(token []byte, powered usbStatus, batteryCapacity uint8, batteryVoltage uint16) error {
	udpHost, err := net.ResolveUDPAddr("udp", currentConfig.Host+":"+strconv.Itoa(currentConfig.Port))
	if err != nil {
		return err
	}

	// 1 byte for battery capacity and power state
	// 2 bytes for battery voltage and power state error flag
	// 8 bytes for timestamp
	const messageLength = 1 + 2 + 8
	message := make([]byte, messageLength)

	// first byte:
	// * first bit indicates if usb power is present
	// * remaining bit indicate battery %
	message[0] = batteryCapacity
	if powered == usbStatusPresent {
		message[0] = message[0] | 0x80
	} else {
		message[0] = message[0] & 0x7F
	}

	// second and third bytes (big endian):
	// * first bit indicates if there was an error reading the usb power state
	// * second and third bits are reserved for future use
	// * remaining bits indicate battery voltage
	binary.BigEndian.PutUint16(message[1:], batteryVoltage)
	if powered == usbStatusError {
		message[1] = message[1] | 0x80
	} else {
		message[1] = message[1] & 0x7F
	}

	// fourth through thirteenth bytes (big endian): local timestamp
	// (to protect against replay attacks)
	binary.BigEndian.PutUint64(message[3:], uint64(time.Now().Unix()))

	mac, err := t.calculateHMAC(message)
	if err != nil {
		return err
	}
	packet := append(message, mac...)

	conn, err := net.DialUDP("udp", nil, udpHost)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Write(packet)
	if err != nil {
		return err
	}

	return nil
}
