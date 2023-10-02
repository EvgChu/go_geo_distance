package yageoservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
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
	url := fmt.Sprintf("%s?format=json&apikey=%s&geocode=%s", urlYandexAPI, os.Getenv(apiKeyEnvVar), address)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func parseYandexAPIResponse(r *http.Response) (*YandexAPIResponse, error) {
	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var result YandexAPIResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func Coordinates(address string) (*YandexAPIResponse, error) {
	resp, err := makeRequest(address)
	log.Infof("Coordinates %#v", resp)
	if err != nil {
		return nil, err
	}
	return parseYandexAPIResponse(resp)
}
