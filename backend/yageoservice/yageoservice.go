package yageoservice

import (
	"fmt"
	"net/http"
	"os"
)

type YandexAPIResponse struct {
	Latitude  float64
	Longitude float64
	Addr      string
}

func (response *YandexAPIResponse) String() string {
	return fmt.Sprintf("Latitude: %f\nLongitude: %f\nAddress: %s", response.Latitude, response.Longitude, response.Addr)
}

const (
	urlYandexAPI = "https://geocode-maps.yandex.ru/1.x"
	apiKeyEnvVar = "YANDEX_API_KEY"
)

func makeRequest(address string) (*http.Response, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?format=json&apikey=%s&geocode=%s", urlYandexAPI, os.Getenv(apiKeyEnvVar), address), nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

type YaGeoService struct{ apiKey string }
