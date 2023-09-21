package yageoservice_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	resp, err := makeRequest(address)
	if err != nil {
		t.Errorf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	var result YandexAPIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatalf("Error unmarshalling JSON response: %v\nBody: %s", err, body)
	}

	if result.Addr != "ул.Тверская, 13, Москва, Россия, 125009" {
		t.Errorf("Invalid address: expected %s, got %s", "ул. Тверская, 13, Москва, Россия", result.Addr)
	}
}
