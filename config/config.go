package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Secret string
}

var Configs Config

func ReadConfig() {
	viper.SetConfigFile(".env") // Tentukan file config
	viper.AutomaticEnv()        // Ambil variabel dari ENV jika ada

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error reading config file: %s", err.Error())
	}

	if err := viper.Unmarshal(&Configs); err != nil {
		fmt.Printf("unable to decode config into struct: %s", err.Error())
	}

}
