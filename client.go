package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var client = &http.Client{}

func makePostRequest[T any, R any](url string, data T) (R, error) {
	var result R

	log.Println("Starting makePostRequest")

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error creating JSON: %v", err)
		return result, errors.New("error creating JSON")
	}
	log.Println("JSON marshaled successfully")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return result, errors.New("error creating request")
	}
	req.Header.Set("Content-Type", "application/json")
	log.Println("Request created successfully")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return result, errors.New("error making request")
	}
	defer resp.Body.Close()
	log.Println("Request made successfully")

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Printf("Error reading response: %v", err)
		return result, errors.New("error reading response")
	}
	log.Println("Response read successfully")

	return result, nil
}
