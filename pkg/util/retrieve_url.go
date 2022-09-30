package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

func RetrieveURL(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("error retrieving site `%s`: %s", url, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error reading response body: %s", err)
	}

	return body
}
