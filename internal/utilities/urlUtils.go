package utilities

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
)

//Placeholder to keep short url key and originalurl in key-value pair
var urlList map[string]string

//It initializes urlList which stores short url key and originalurl in key-value pair.
func init() {
	urlList = map[string]string{}
}

//GetUrlList provides access to urlList outside utilities package.
func GetUrlList() map[string]string {
	return urlList
}

// GetShortURL validates the requested url and generates a shorter url
//
// It sends a http response with short url if the request url is valid
// Otherwise it sends a http BadRequest(Status code 400) as response
func GetShortURL(w http.ResponseWriter, r *http.Request) {
	reqURL, ok := r.URL.Query()["url"]
	log.Debug("Requested url is : " + reqURL[0])

	//check if requested URL is valid
	if ok {

		if isValidURL(reqURL[0]) {
			shortUrl := GenerateURL(reqURL[0])
			fmt.Fprintf(w, shortUrl)
			return

		}
	}

	//If the code execution reached here that means there is error in the request
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "Not a valid url")
}

//isValidURL checks for the regular expression of the given url
//url should start with http or https
//url should be in the following pattern [http/https]://xyz.com
//
//if the requested url is a valid url then it returns true
//if the requested url is not a valid url then it returns false
func isValidURL(reqUrl string) bool {
	_, err := url.ParseRequestURI(reqUrl)

	if err == nil {
		return true
	} else {
		return false
	}

}

//GenerateURL generates a unique key according to the url string passed
//
//It returns the short url in the format http://localhost:8080/shorturl/<key>
func GenerateURL(reqUrl string) string {
	genkey := GenerateKey(reqUrl)
	urlList[genkey] = reqUrl
	return Data["appconfig"].AppURL + "shorturl/" + genkey
}

//RedirectToOriginalUrl gets invoked when short url is visited and it redirects to the original link
//
//If the short url visited does not exist in the urlList then it sends a http BadRequest(Status code 400) as response
func RedirectToOriginalUrl(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pathArgs := strings.Split(path, "/")
	if val, ok := urlList[pathArgs[2]]; ok {
		log.Info("Redirected to: " + val)
		http.Redirect(w, r, val, http.StatusPermanentRedirect)
	} else {
		log.Warn("Invalid request URL")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Not a valid url")

	}
}
