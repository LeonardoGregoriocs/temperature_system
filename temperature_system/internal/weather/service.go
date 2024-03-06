package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

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
	escapedCityName := url.QueryEscape(city)
	URL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%v&q=%v&aqi=no", apiKey, escapedCityName)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Erro ao criar a requisição:", err)
		return &Weather{}, err
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err)
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
