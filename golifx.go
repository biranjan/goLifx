package golifx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	bodyString string
)

// Lights of status
type Lights struct {
	Label     string `json:"label"`
	Power     bool   `json:"power"`
	Connected bool   `json:"connected"`
}

// Declare result array
type Results struct {
	Results []Result `json:"results"`
}

// Rseults values
type Result struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Label  string `json:"label"`
}

func queryapi(url string, reqtype string, body string, listLight bool) (string, string) {

	//newbody := strings.NewReader(body)
	//newbody := strings.NewReader(`power=` + power)
	var statmessage string
	var light []Lights
	var results Results
	req, err := http.NewRequest(reqtype, url, strings.NewReader(body))
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer <insert your api key here>")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatal(err)
	}
	//fmt.Println("HTTP Response Status: " + strconv.Itoa(resp.StatusCode))
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		statmessage = "request succeded "
		// initialize our array
		if listLight == true {
			json.Unmarshal([]byte(bodyBytes), &light)
			for i := range light {
				fmt.Printf("label: %s, connected: %s, power: %s  \n", light[i].Label, strconv.FormatBool(light[i].Power), strconv.FormatBool(light[i].Connected))
			}
			bodyString = ""
		} else {

			bodyString = string(bodyBytes)
		}
	} else if resp.StatusCode == 207 {

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		statmessage = "request succeded & status code 207"
		// initialize our array
		json.Unmarshal([]byte(bodyBytes), &results)

		for i := range results.Results {

			fmt.Printf("label: %s, connected: %s, \n", results.Results[i].Label, results.Results[i].Status)
		}
		bodyString = ""
	} else {
		statmessage = "request failed: Make sure sure light is connected & try again" // if you want to see stat code do <strconv.Itoa(resp.StatusCode)> instead
	}
	return bodyString, statmessage
}

