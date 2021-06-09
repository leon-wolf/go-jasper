package gojasper

import (
	"io/ioutil"
	"net/http"
)

func (a API) ServerInfo() ([]byte, error) {
	// Request (GET https://localhost/jasperserver/rest_v2/serverInfo)

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", a.URL+"/serverInfo", nil)
	if err != nil {
		return nil, err
	}

	// Headers
	req.Header.Add("Accept", "application/json")

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
