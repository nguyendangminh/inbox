package fb

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpget(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Sending HTTP request failed:", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Reading response's body failed:", err.Error())
		return nil, err
	}

	if resp.StatusCode != 200 { // Just handle success case
		e := errors.New(fmt.Sprintf("status code %d - %s", resp.StatusCode, string(body)))
		log.Println("Request failed:", e.Error())
		return nil, e
	}

	return body, nil
}