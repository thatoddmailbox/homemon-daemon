package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"
	"net"
	"strconv"
)

const port = 9325

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
	log.Println(currentConfig.Host + ":" + strconv.Itoa(port))
	udpHost, err := net.ResolveUDPAddr("udp", currentConfig.Host+":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	message := []byte("dab dab dab")
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
