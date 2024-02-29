package weather

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/leonardogregoriocs/temperature_system/temperature_system/internal/utils"
	"github.com/spf13/viper"
)

type Configs struct {
	ApiKey string `json:"ApiKey"`
}

func GetWeatherByCity(city string) (*Weather, error) {
	b, err := json.Marshal(viper.Get("ApiKey"))
	if err != nil {
		panic(err)
	}

	apiKey := utils.RemoveFirstAndLast(b)
	resp, err := http.Get("http://api.weatherapi.com/v1/current.json?key=" + apiKey + "&q=" + city + "&aqi=no")

	if err != nil {
		return &Weather{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != err {
		return &Weather{}, err
	}

	var weatherData WeatherDataPayload
	err = json.Unmarshal(body, &weatherData)

	if err != nil {
		return &Weather{}, err
	}

	weatherObj := conversionWeather(weatherData.Current.TemperatureC, weatherData.Current.TemperatureF)

	return &weatherObj, nil

}

func conversionWeather(celsius, fahrenheit float64) Weather {
	return Weather{
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
		Kelvin:     celsius + 273.15,
	}
}
