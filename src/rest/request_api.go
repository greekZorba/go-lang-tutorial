package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
const (
	apiUrl = "https://korean.visitkorea.or.kr/list/ms_list.do?areacode=2"
)

func main() {
	request, err := http.NewRequest("GET", apiUrl, nil)
	errorCheck(err)

	authorizationKey := ""
	request.Header.Add("Authorization", authorizationKey)
	request.Header.Add("Content-Type", "application/json")

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
