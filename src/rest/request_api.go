package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiUrl = "https://openapi.leisureq.com/deals/{id}"
)

func main() {
	request, err := http.NewRequest("GET", apiUrl, nil)
	errorCheck(err)

	authorizationKey := ""
	request.Header.Add("x-api-key", authorizationKey)
	request.Header.Add("x-api-version", "v2")
	request.Header.Add("Content-Type", "application/json")

	request.URL.Path = strings.Replace(request.URL.Path, "{id}", "56892", 1)

	queryParam, err := url.ParseQuery(request.URL.RawQuery)
	queryParam.Add("clientId", "")
	request.URL.RawQuery = queryParam.Encode()

	client := &http.Client{}
	response, err := client.Do(request)
	errorCheck(err)

	body, _ := ioutil.ReadAll(response.Body)
	log.Println(string(body))

	// @see https://stackoverflow.com/questions/51452148/how-can-i-make-a-request-with-a-bearer-token-in-go

}

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}
