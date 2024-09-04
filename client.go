package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var client = &http.Client{}

func makePostRequest[T any, R any](url string, data T) (R, error) {
	var result R

	jsonData, err := json.Marshal(data)
	if err != nil {
		return result, errors.New("error creating JSON")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return result, errors.New("error creating request")
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return result, errors.New("error making request")
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return result, errors.New("error reading response")
	}

	return result, nil
}
