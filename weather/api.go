package weather

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func CurrentTemperature(providerUrl string, city string) (int, error) {
	rawResponse, err := http.Get(providerUrl + "/current?query=" + city)
	if err != nil {
		return -90, errors.New("could not contact api")
	}
	responseBody, err := ioutil.ReadAll(rawResponse.Body)
	if err != nil {
		return 0, err
	}
	var response Response
	if err := json.Unmarshal(responseBody, &response); err != nil {
		return -90, errors.New("could not decode weather")
	}
	return response.Current.Temperature, nil
}

type Response struct {
	Current struct {
		Temperature int `json:"temparature"`
	} `json:"current"`
}
