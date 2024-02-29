package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/leonardogregoriocs/temperature_system/temperature_system/internal/handler"
	"github.com/spf13/viper"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/weather", handler.WeatherHandler)

	fmt.Println("Starting web server on port: 8080")

	http.ListenAndServe(":8080", r)
}

func init() {
	viper.SetConfigName("configs")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../..")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
