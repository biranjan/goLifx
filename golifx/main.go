package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/biranjan/golifx"
)

// flag
var (
	user       string
	power      string
	list       string
	label      string
	bodyString string
	url        string
)

func init() {
	flag.StringVar(&power, "power", "na", "change power state")
	flag.StringVar(&list, "list", "na", "change power state")
	flag.StringVar(&label, "label", "na", "change power state of light with particular label")
}

func main() {
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
		status, message := golifx.queryapi(url, reqtype, body, false)
		print(message)
		print(status)

	}

	if list != "na" {
		body := `nil`
		url := "https://api.lifx.com/v1/lights/all"
		reqtype := "GET"
		status, message := golifx.queryapi(url, reqtype, body, true)
		print(message)
		print(status)

	}

}
