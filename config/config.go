package config

import (
	"fmt"
	"notebook/storage"

	"github.com/spf13/viper"
)

type Config struct {
	Storage string `mapstructure:"storage"`
	Email   string `mapstructure:"email"`
}

var config Config

func Configure() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Storage)
	fmt.Println(config.Email)

	storage.Initialize(config.Storage)
}
