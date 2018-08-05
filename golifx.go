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

/* // flag
var (
	user       string
	power      string
	list       string
	label      string
	bodyString string
	url        string
) */

var (
	bodyString string
)

// Lights of lights
type Lights struct {
	Label     string `json:"label"`
	Power     bool   `json:"power"`
	Connected bool   `json:"connected"`
}

// Users struct which contains
// an array of users
type Results struct {
	Results []Result `json:"results"`
}

// User struct which contains a name
// a type and a list of social links
type Result struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Label  string `json:"label"`
}

/* func init() {
	flag.StringVar(&power, "power", "na", "change power state")
	flag.StringVar(&list, "list", "na", "change power state")
	flag.StringVar(&label, "label", "na", "change power state of light with particular label")
} */

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

/* func main() {
	flag.Parse()
	// if user does not supply flags, print usage
	// we can clean this up later by putting this into its own function
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	//users := strings.Split(user, ",")
	//fmt.Printf("Searching user(s): %s\n", user)

	if power != "na" {
		body := (`power=` + power)
		reqtype := "PUT"
		if label != "na" {
			url = fmt.Sprintf("https://api.lifx.com/v1/lights/label:%s/state", label)
		} else {
			url = "https://api.lifx.com/v1/lights/all/state"
		}
		status, message := queryapi(url, reqtype, body, false)
		print(message)
		print(status)

	}

	if list != "na" {
		body := `nil`
		url := "https://api.lifx.com/v1/lights/all"
		reqtype := "GET"
		status, message := queryapi(url, reqtype, body, true)
		print(message)
		print(status)

	}

}
*/
