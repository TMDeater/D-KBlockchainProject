package util

import (
	// Add Golang imports here
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	// Add Hyperledger imports here
	// Add 3rd part imports here
	// Add local imports here
)

// PostJSON Utility function post a JSON format string `body` to address `url`
func PostJSON(url string, json []byte) ([]byte, error) {

	client := http.Client{}

	// Prepare request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("non-success status code %d returned", response.StatusCode)
	}

	// Return response body
	return ioutil.ReadAll(response.Body)
}
