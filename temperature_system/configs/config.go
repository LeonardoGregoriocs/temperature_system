package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type conf struct {
	ApiKey string `mapstructure:"API_KEY"`
}

func LoadConfig() (*conf, error) {
	var cfg *conf
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
