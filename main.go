package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

	req, err := http.NewRequest("POST", reportURL+"?p=1&b=50&v=4000", nil)
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
