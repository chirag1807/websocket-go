package config

import (
	"github.com/chirag1807/websocket-go/api/model/dto"
	"fmt"

	"github.com/spf13/viper"
)

var Config dto.Config

func LoadEnv(envFilePath string) {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(envFilePath)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Println("Error While Decoding .env File")
	}

}
