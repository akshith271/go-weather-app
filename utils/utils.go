package utils

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/akshith271/weather/models.go"
)

const ApiURL = "http://api.openweathermap.org/data/2.5/weather?APPID="

func loadApiConfig(fileName string) (models.ApiConfigData, error) {
	// helps us get the api key from the .apiConfig file
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return models.ApiConfigData{}, err
	}
	var c models.ApiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return models.ApiConfigData{}, err
	}
	return c, nil
}

var Query = func(city string) (models.WeatherData, error) {
	apiConfig, err := loadApiConfig("apiConfig/.apiConfig")
	if err != nil {
		return models.WeatherData{}, err
	}
	resp, err := http.Get(ApiURL + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return models.WeatherData{}, err
	}
	defer resp.Body.Close()
	var data models.WeatherData
	if json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return models.WeatherData{}, err
	}
	return data, nil
}
