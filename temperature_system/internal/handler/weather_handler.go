package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/leonardogregoriocs/temperature_system/temperature_system/internal/weather"
	"github.com/leonardogregoriocs/temperature_system/temperature_system/internal/zipcode"
)

var (
	ErrInvalidZipcode = ("invalid zipcode")
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var payload Payload
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	dataCep, err := zipcode.GetAddressViaCep(payload.Zipcode)
	if err != nil {
		switch {
		case err.Error() == "invalid zipcode":
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		case err.Error() == "can not find zipcode":
			http.Error(w, "invalid zipcode", http.StatusNotFound)
			return
		}
	}

	result, err := weather.GetWeatherByCity(dataCep.Localidade)
	if err != nil {
		http.Error(w, "Error get weather by city", http.StatusInternalServerError)
		return
	}

	resultJson, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error json marshal", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resultJson)
}
