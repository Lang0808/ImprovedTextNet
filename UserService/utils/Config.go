package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() (err error) {
	viper.SetConfigFile("config/app.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
