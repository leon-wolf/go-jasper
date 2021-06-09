package gojasper

import (
	"io/ioutil"
	"net/http"
)

func (a API) RunReport(path, format string, params ...interface{}) ([]byte, error) {

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("GET", a.URL+"/reports"+path, nil)
	if err != nil {
		return nil, err
	}

	// Headers
	req.Header.Add("Accept", "application/json")

	// Execute Report
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
