package main

import (
	"log"
	"net/http"
	"urlapp/initapp"
	"urlapp/internal/utilities"
)

var serverPort string

//init initializes all the configurations and sets serverport as per the settings in config.yaml
func init() {
	initapp.Initialize()
	serverPort = utilities.Data["appconfig"].AppPort
}

//Entry point of the Application: main function
func main() {

	http.HandleFunc("/shorten", utilities.GetShortURL)
	http.HandleFunc("/shorturl/", utilities.RedirectToOriginalUrl)

	log.Fatal(http.ListenAndServe(":"+serverPort, nil))

}
