package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type transportHTTP struct{}

func (t *transportHTTP) Transport(token []byte, powered usbStatus, batteryCapacity uint8, batteryVoltage uint16) error {
	usbPresentString := "0"
	if powered == usbStatusError {
		usbPresentString = "-1"
	} else if powered == usbStatusPresent {
		usbPresentString = "1"
	}

	params := url.Values{
		"p": []string{usbPresentString},
		"b": []string{strconv.Itoa(int(batteryCapacity))},
		"v": []string{strconv.Itoa(int(batteryVoltage))},
	}

	req, err := http.NewRequest("POST", reportURL+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Token", string(token))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	responseData := map[string]interface{}{}
	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	return nil
}
