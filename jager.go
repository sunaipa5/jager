package jager

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func IsJSON(data []byte) bool {
	return json.Valid(data)
}

// Convert struct to json and write response
func Struct(w http.ResponseWriter, jsonData interface{}) error {
	encodedJson, err := json.Marshal(jsonData)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encodedJson)
	return nil
}

func Write(w http.ResponseWriter, jsonData []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// Reading JSON data in the Http request body
func Read(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("request body empty")
	}

	jsonCheck := IsJSON(body)
	if jsonCheck {
		return body, nil
	}
	return nil, fmt.Errorf("request body not in JSON format")
}

// Convert string to json and write response
/* Sample:

jager.String(w,`{"name":"john","surname":"doe"}`)

*/
func String(w http.ResponseWriter, input string) error {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(input), &jsonMap)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(jsonMap)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	return nil
}

// Create a json with map and write response
/* Sample:

jager.Map(w, map[string]interface{}{
	"name":      "John",
	"surname":   "Doe",
	"age":       30,
	"isStudent": true,
})

*/
func Map(w http.ResponseWriter, data map[string]interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

	return nil
}
