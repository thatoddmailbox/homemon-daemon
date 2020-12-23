package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const baseURL = "https://homemon-rpt.studer.dev/"
const reportURL = baseURL + "report"

func report() error {
	tokenBytes, err := ioutil.ReadFile("token.txt")
	if err != nil {
		return err
	}
	tokenString := strings.TrimSpace(string(tokenBytes))

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

	params := url.Values{
		"p": []string{"1"},
		"b": []string{strconv.Itoa(batteryCapacity)},
		"v": []string{strconv.Itoa(batteryVoltage)},
	}

	req, err := http.NewRequest("POST", reportURL+"?"+params.Encode(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Token", tokenString)
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
	log.Println(responseData)
	return nil
}

func main() {
	log.Println("homemon-daemon")

	err := report()
	if err != nil {
		panic(err)
	}
}
